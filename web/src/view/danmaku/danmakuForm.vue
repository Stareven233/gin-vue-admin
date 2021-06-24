<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="日期:">
    
      <el-date-picker type="date" placeholder="选择日期" v-model="formData.date" clearable></el-date-picker>
    </el-form-item>
    
      <el-form-item label="视频时间:">
    
      <el-input-number v-model="formData.time" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="内容:">
    <el-input v-model="formData.text" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="颜色:">
    <el-input v-model.number="formData.color" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="类型:">
    <el-input v-model.number="formData.type" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="发布者id:">
    <el-input v-model.number="formData.authorId" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="视频id:">
    <el-input v-model.number="formData.videoId" clearable placeholder="请输入"/>
    </el-form-item>
    <el-form-item>
        <el-button type="primary" @click="save">保存</el-button>
        <el-button type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  createDanmaku,
  updateDanmaku,
  findDanmaku
} from '@/api/danmaku' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Danmaku',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            date: new Date(),
            time: 0,
            text: '',
            color: 0,
            type: 0,
            authorId: 0,
            videoId: 0,
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findDanmaku({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.redanmaku
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createDanmaku(this.formData)
          break
        case 'update':
          res = await updateDanmaku(this.formData)
          break
        default:
          res = await createDanmaku(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>