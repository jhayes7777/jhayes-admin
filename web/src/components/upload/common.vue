<template>
  <div>
    <el-upload
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
      :before-upload="checkFile"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      :data="{ classId: props.classId }"
      :headers="{ 'x-token': token }"
      multiple
      class="upload-btn"
    >
      <el-button type="primary" :icon="Upload">普通上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { isVideoMime, isImageMime } from '@/utils/image'
  import { getBaseUrl } from '@/utils/format'
  import { Upload } from '@element-plus/icons-vue'
  import { useUserStore } from '@/pinia'

  defineOptions({
    name: 'UploadCommon'
  })

  const userStore = useUserStore()

  const token = userStore.token

  const props = defineProps({
    classId: {
      type: Number,
      default: 0
    }
  })

  const emit = defineEmits(['on-success'])

  const fullscreenLoading = ref(false)

  const checkFile = (file) => {
    console.log(
      file,
      file.type,
      'file.typefile.typefile.typefile.typefile.typefile.typefile.type'
    )
    fullscreenLoading.value = true
    const isLt1M = file.size / 1024 / 1024 < 1 // 500K, @todo 应支持在项目中设置
    const isLt50M = file.size / 1024 / 1024 < 50 // 5MB, @todo 应支持项目中设置
    const isLt100M = file.size / 1024 / 1024 < 100 // 5MB, @todo 应支持项目中设置

    const isVideo = isVideoMime(file.type)
    const isImage = isImageMime(file.type)
    const isModel =
      file.name.includes('.glb') ||
      file.name.includes('.GLB') ||
      file.name.includes('.fbx') ||
      file.name.includes('.FBX')
    let pass = true
    if (!isVideo && !isImage && !isModel) {
      ElMessage.error(
        '上传图片只能是 jpg,png,svg,webp 格式, 上传视频只能是 mp4,webm 格式, 上传模型只能是 glb,fbx格式!'
      )
      fullscreenLoading.value = false
      pass = false
    }
    if (!isLt100M && isModel) {
      ElMessage.error('上传模型大小不能超过 100MB')
      fullscreenLoading.value = false
      pass = false
    }
    if (!isLt50M && isVideo) {
      ElMessage.error('上传视频大小不能超过 50MB')
      fullscreenLoading.value = false
      pass = false
    }
    if (!isLt1M && isImage) {
      ElMessage.error('未压缩的上传图片大小不能超过 1MB，请使用压缩上传')
      fullscreenLoading.value = false
      pass = false
    }

    console.log('upload file check result: ', pass)

    return pass
  }

  const uploadSuccess = (res) => {
    const { data } = res
    if (data.file) {
      emit('on-success', data.file.url)
    }
  }

  const uploadError = () => {
    ElMessage({
      type: 'error',
      message: '上传失败'
    })
    fullscreenLoading.value = false
  }
</script>
