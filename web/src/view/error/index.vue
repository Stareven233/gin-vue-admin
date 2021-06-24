<template>
  <div>
    <div class="big">
      <div class="inner">
        <img src="../../assets/notFound.png">
        <p>这个页面不存在于这个次元</p>
        <p style="font-size:18px;line-height:40px;">可能是菜单未分配仪表盘</p>
        <p>↓</p>
        <img src="../../assets/Furedorika.png" class="leftPic">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Error'

}
</script>

<style lang="scss">
    .big{
        width: 100%;
        height: calc(100vh - 220px);
        background-color: rgb(244,244,244);
        position: relative;
    }
    .inner{
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    .inner p {
        text-align: center;
        font-size: 24px;
    }
    .inner .leftPic{
        width: 60px;
        height: 60px;
        margin-left: 44%;
        margin-top: 20px;
    }
</style>
