<template>
  <div class="absolute w-full h-full">
    <el-button
      @click="changeModel"
      title="随机模型"
      class="bg-[#00000000] border-[#ffffff00] absolute right-[6rem] z-[999] w-[10rem] h-[10rem] bottom-[6rem] bg-cover hover:animate-shake cursor-pointer bg-[url('@/assets/model.png')] bg-no-repeat bg-center"
    ></el-button>
    <div id="my-three" ref="container" class=""></div>
  </div>
</template>
<script setup>
  import {
    getFileList,
    deleteFile,
    editFileName,
    importURL
  } from '@/api/fileUploadAndDownload'
  import {
    getCurrentInstance,
    ref,
    onMounted,
    onUnmounted,
    reactive
  } from 'vue'

  import * as THREE from 'three'
  import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls'
  import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader' //gltf
  import { FontLoader } from 'three/examples/jsm/loaders/FontLoader'
  import { TextGeometry } from 'three/addons/geometries/TextGeometry'
  // 导入后处理相关模块
  import { EffectComposer } from 'three/addons/postprocessing/EffectComposer.js'
  import { RenderPass } from 'three/addons/postprocessing/RenderPass.js'
  import { UnrealBloomPass } from 'three/addons/postprocessing/UnrealBloomPass.js'
  import { ShaderPass } from 'three/addons/postprocessing/ShaderPass.js'
  import { CopyShader } from 'three/addons/shaders/CopyShader.js'
  import bgImg from '@/assets/c83e0d6cbeb34c18b39622647d5bd2ab.jpeg'
  const { proxy } = getCurrentInstance()
  const container = ref(null)
  // 存储Three.js相关对象
  let scene,
    camera,
    renderer,
    controls,
    loader,
    model,
    ringMesh,
    textureLoader,
    composer

  // 状态管理
  const state = reactive({
    modelUrl: 'https://threejs.org/examples/models/gltf/LittlestTokyo.glb',
    isModelLoaded: false,
    ringRotationSpeed: 0.003,
    ringRadius: 30,
    ringWidth: 200,
    runeDensity: 24,
    bloomStrength: 1.5,
    bloomThreshold: 0,
    bloomRadius: 0.8,
    backgroundColor: 0x11111100, // 背景颜色（用于非透明模式）
    backgroundOpacity: 0.0 // 背景透明度（0=完全透明，1=完全不透明）
  })
  const tableData = ref([])
  const total = ref(0)
  onMounted(() => {
    getModelData().then(() => {
      init()
      renderModel()
      gltfModel1(
        'http://localhost:8080/api/' +
          tableData.value[Math.floor(Math.random() * tableData.value.length)]
      )
      // 创建辉光圆环
      createGlowRing()

      // 设置后处理
      setupPostProcessing()
      render()
    })
  })

  // 之后需要删除模型时
  function removeModel(mesh) {
    // 从场景中移除网格
    if (mesh.parent) {
      mesh.parent.remove(mesh)
    }

    // 释放几何体资源
    if (mesh.geometry) {
      mesh.geometry.dispose()
    }

    // 释放材质资源（如果多个网格共享同一材质，需谨慎处理）
    if (mesh.material) {
      // 单个材质
      if (!Array.isArray(mesh.material)) {
        mesh.material.dispose()
      }
      // 多个材质组成的数组
      else {
        mesh.material.forEach((material) => material.dispose())
      }
    }

    // 清除引用，帮助垃圾回收
    mesh = null
  }

  function changeModel() {
    // 复制场景中的子对象数组（避免在循环中修改原数组）
    console.log(scene.children)
    const children = [...scene.children]

    children.forEach((child) => {
      // 只处理网格对象（Mesh、SkinnedMesh等）
      if (!child.isLight) {
        removeModel(child) // 使用上面定义的removeModel函数
      }
    })
    gltfModel1(
      'http://localhost:8080/api/' +
        tableData.value[Math.floor(Math.random() * tableData.value.length)]
    )
    createGlowRing()

  }
  // 清理函数
  const cleanup = () => {
    if (renderer && container && container.value) {
      renderer.dispose()
      container.value.removeChild(renderer.domElement)
    }

    if (controls) controls.dispose()

    if (composer) {
      composer.renderTarget1.dispose()
      composer.renderTarget2.dispose()
    }

    window.onresize = function () {}
  }
  onUnmounted(() => {
    cleanup()
  })
  async function getModelData() {
    const table = await getFileList({
      classId: 1,
      page: 1,
      pageSize: 1000
    })
    if (table.code === 0) {
      tableData.value = table.data.list.map((x) => x.url)
      total.value = table.data.total
    }
  }
  function init() {
    if (!container.value) return
    scene = new THREE.Scene()
    renderer = new THREE.WebGLRenderer({
      //抗锯齿属性，WebGLRenderer常用的一个属性
      antialias: true,
      alpha: true
    })
    // renderer.setClearAlpha(0)

    loader = new GLTFLoader()
    camera = new THREE.PerspectiveCamera(
      45,
      container.value.clientWidth / container.value.clientHeight,
      1,
      100000
    )
    controls = new OrbitControls(camera, renderer.domElement) // 添加光源
    const ambientLight = new THREE.AmbientLight(0x696969, 2)
    scene.add(ambientLight)

    const directionalLight = new THREE.DirectionalLight(0xffffff, 2)
    directionalLight.position.set(5, 10, 7)
    directionalLight.castShadow = true
    scene.add(directionalLight)
    scene.background = new THREE.Color('#262626') // 白色背景

    //设置相机位置
    camera.position.set(0, 80, 250)
    //设置相机方向
    camera.lookAt(0, 0, 0)

    renderer.setSize(container.value.clientWidth, container.value.clientHeight)
    renderer.setPixelRatio(window.devicePixelRatio)
    renderer.antialias = true
    container.value.appendChild(renderer.domElement)
  }

  function gltfModel1(modelUrl) {
    if (!container.value || !renderer || !camera || !controls || !loader) return

    // glb
    loader.load(
      modelUrl,
      function (gltf) {
        gltf.scene.scale.x = 100
        gltf.scene.scale.y = 100
        gltf.scene.scale.z = 100
        scene.add(gltf.scene)
        model = gltf.scene

        model.traverse((child) => {
          if (child.isMesh) {
            // child.castShadow = true
            // child.receiveShadow = true
            child.userData.isGlowing = false
            // if (
            //   child.name.includes('building') ||
            //   child.name.includes('Tower')
            // ) {
            //   child.material.emissive = new THREE.Color(0x224488)
            //   child.material.emissiveIntensity = 0.5
            // }
          }
        })
        // 计算模型边界以确定圆环大小
        const box = new THREE.Box3().setFromObject(gltf.scene)
        const center = box.getCenter(new THREE.Vector3())
        const size = box.getSize(new THREE.Vector3())
        const maxDimension = Math.max(size.x, size.y, size.z)
      },
      function (xhr) {
        const percent = Math.floor((xhr.loaded / xhr.total) * 100) // 计算加载进度百分比
      }
    )
  }

  function renderModel() {
    if (!container.value || !renderer || !camera || !controls || !loader) return

    //渲染
    renderer.setSize(container.value.clientWidth, container.value.clientHeight) //设置渲染区尺寸
    renderer.render(scene, camera) //执行渲染操作、指定场景、相机作为参数
    // renderer.setClearColor(0x00ff00); // 设置背景颜色为绿色
    renderer.toneMapping = THREE.ACESFilmicToneMapping
    // 设置曝光度
    renderer.toneMappingExposure = 1.5 // 适当调整曝光度

    controls.minPolarAngle = Math.PI / 4 // 最小极角为 45 度
    controls.maxPolarAngle = Math.PI / 2 // 最大极角为 90 度
  }

  function render() {
    // 旋转圆环
    if (ringMesh) {
      ringMesh.rotation.z += state.ringRotationSpeed
    }

    // 使用后处理器渲染场景
    if (composer) {
      composer.render()
    } else {
      renderer.render(scene, camera)
    }
    controls.update()
    requestAnimationFrame(render)
  }

  // 画布跟随窗口变化
  window.onresize = function () {
    if (!container.value || !renderer || !camera || !controls || !loader) return
    renderer.setSize(container.value.clientWidth, container.value.clientHeight)
    camera.aspect = container.value.clientWidth / container.value.clientHeight
    camera.updateProjectionMatrix()
    // 更新后处理尺寸
    if (composer) {
      composer.setSize(
        container.value.clientWidth,
        container.value.clientHeight
      )
    }
  }

  // 交互事件
  addEventListener('dblclick', onMouseDblclick, false)
  function onMouseDblclick(event) {
    let intersects = getIntersects(event)

    if (intersects.length !== 0 && intersects[0].object instanceof THREE.Mesh) {
      const selectedObject = intersects[0].object
      let selectedObjects = []
      selectedObjects.push(selectedObject)
      console.log(selectedObjects)
      // outlinePass.selectedObjects = selectedObjects;
    }
  }
  //获取与射线相交的对象数组
  function getIntersects(event) {
    let rayCaster = new THREE.Raycaster()
    let mouse = new THREE.Vector2()

    //通过鼠标点击位置，计算出raycaster所需点的位置，以屏幕为中心点，范围-1到1
    mouse.x = (event.clientX / window.innerWidth) * 2 - 1
    mouse.y = -(event.clientY / window.innerHeight) * 2 + 1 //这里为什么是-号，没有就无法点中

    //通过鼠标点击的位置(二维坐标)和当前相机的矩阵计算出射线位置
    rayCaster.setFromCamera(mouse, camera)
    return rayCaster.intersectObjects(scene.children)
  }

  // 创建符文纹理
  const createRuneTexture = () => {
    // 创建画布
    const canvas = document.createElement('canvas')
    const size = 256
    canvas.width = size
    canvas.height = size

    const ctx = canvas.getContext('2d')

    // 清空画布
    ctx.fillStyle = 'rgba(0, 0, 0, 0)'
    ctx.fillRect(0, 0, size, size)

    // 设置符文样式
    ctx.fillStyle = 'rgba(255, 255, 255, 0.8)'
    ctx.font = 'bold 24px Arial'
    ctx.textAlign = 'center'
    ctx.textBaseline = 'middle'

    // 绘制符文（使用随机符号模拟）
    const runes = ['Ψ', 'Φ', 'Ω', 'Σ', 'Λ', 'Θ', 'Ξ', 'Δ', 'Γ', 'Ξ', 'Π', 'Ψ']
    const center = size / 2
    const radius = size / 2.5

    for (let i = 0; i < state.runeDensity; i++) {
      const angle = (i / state.runeDensity) * Math.PI * 2
      const x = center + Math.cos(angle) * radius
      const y = center + Math.sin(angle) * radius

      // 随机选择符文
      const rune = runes[Math.floor(Math.random() * runes.length)]

      // 旋转符文
      ctx.save()
      ctx.translate(x, y)
      ctx.rotate(angle + Math.PI / 2)
      ctx.fillText(rune, 0, 0)
      ctx.restore()
    }

    // 创建Three.js纹理
    const texture = new THREE.CanvasTexture(canvas)
    texture.wrapS = THREE.RepeatWrapping
    texture.wrapT = THREE.RepeatWrapping
    texture.repeat.set(1, 1)

    return texture
  }

  // 创建辉光圆环
  const createGlowRing = () => {
    // 生成符文纹理
    const runeTexture = createRuneTexture()

    // 创建环形几何体
    const geometry = new THREE.RingGeometry(
      state.ringRadius - state.ringWidth / 2,
      state.ringRadius + state.ringWidth / 2,
      64
    )

    // 创建辉光材质
    const material = new THREE.MeshBasicMaterial({
      map: runeTexture,
      transparent: true,
      side: THREE.DoubleSide,
      color: 0x00ccff,

      opacity: 0.9 // 稍微降低不透明度增强透明效果
    })

    // 创建圆环网格
    ringMesh = new THREE.Mesh(geometry, material)
    ringMesh.rotation.x = Math.PI / 2 // 使其水平放置
    // ringMesh.rotation.z = Math.PI / 2 // 使其水平放置
    // ringMesh.rotation.y = 2*Math.PI / 4.1 // 使其水平放置
    ringMesh.position.y = 0 // 调整高度位置
    // 添加辉光效果（通过后处理实现）
    const glowMaterial = new THREE.MeshBasicMaterial({
      color: 0x00ccff,
      transparent: true,
      opacity: state.glowIntensity,
      side: THREE.DoubleSide,
      depthWrite: false
    })

    // const glowMesh = new THREE.Mesh(geometry, glowMaterial)
    // glowMesh.rotation.x = Math.PI / 2
    // glowMesh.position.y = -2.01 // 稍微偏移以避免z-fighting
    // glowMesh.scale.multiplyScalar(1.05) // 使辉光略大于符文环
    // 为后处理标记辉光对象
    ringMesh.userData.isGlowing = true

    scene.add(ringMesh)
    // scene.add(glowMesh)
  }
  // 设置后处理管线
  const setupPostProcessing = () => {
    // 创建后处理器
    const renderTarget = new THREE.WebGLRenderTarget(
      container.value.clientWidth,
      container.value.clientHeight,
      {
        format: THREE.RGBAFormat, // 关键：使用RGBA格式支持透明度
        alpha: true,
        minFilter: THREE.LinearFilter,
        magFilter: THREE.LinearFilter
      }
    )

    composer = new EffectComposer(renderer, renderTarget)

    // 添加渲染通道
    const renderPass = new RenderPass(scene, camera)
    console.log(renderPass, renderTarget)
    // renderPass.clearAlpha = 0 // 完全透明
    composer.addPass(renderPass)

    // 添加UnrealBloomPass实现辉光效果
    const bloomPass = new UnrealBloomPass(
      new THREE.Vector2(
        container.value.clientWidth,
        container.value.clientHeight
      ),
      state.bloomStrength,
      state.bloomRadius,
      state.bloomThreshold
    )
    composer.addPass(bloomPass)

    // 创建自定义着色器通道以保留透明度
    const finalPass = new ShaderPass({
      uniforms: {
        tDiffuse: { value: null },
        opacity: { value: 1.0 }
      },
      vertexShader: `
      varying vec2 vUv;
      void main() {
        vUv = uv;
        gl_Position = projectionMatrix * modelViewMatrix * vec4( position, 1.0 );
      }
    `,
      fragmentShader: `
      uniform sampler2D tDiffuse;
      uniform float opacity;
      varying vec2 vUv;
      void main() {
        gl_FragColor = texture2D( tDiffuse, vUv );
      }
    `
    })
    finalPass.renderToScreen = true
    // composer.addPass(finalPass)
  }
</script>

<style>
  #my-three {
    position: absolute;
    width: 100%;
    height: 100%;
    background-image: url('@/assets/c83e0d6cbeb34c18b39622647d5bd2ab.jpeg');
    background-repeat: no-repeat;
    background-size: 100% 100%;
  }
</style>
