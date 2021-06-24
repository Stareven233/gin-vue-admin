<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="标题:">
    <el-input v-model="formData.title" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="文件路径:">
    <el-input v-model="formData.file" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="发布日期:">
    
      <el-date-picker type="date" placeholder="选择日期" v-model="formData.date" clearable></el-date-picker>
    </el-form-item>
    
      <el-form-item label="作者id:">
    <el-input v-model.number="formData.authorId" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="类型:">
    <el-input v-model.number="formData.type" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="封面:">
    <el-input v-model="formData.face" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="简介:">
    <el-input v-model="formData.desc" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="硬币数:">
    <el-input v-model.number="formData.coin" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="收藏数:">
    <el-input v-model.number="formData.collect" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="点赞数:">
    <el-input v-model.number="formData.like" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="审核状态:">
    <el-input v-model.number="formData.status" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="审核者:">
    <el-input v-model.number="formData.inspector" clearable placeholder="请输入"/>
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
  createVideo,
  updateVideo,
  findVideo
} from '@/api/video' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Video',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            title: '',
            file: '',
            date: new Date(),
            authorId: 0,
            type: 0,
            face: '',
            desc: '',
            coin: 0,
            collect: 0,
            like: 0,
            status: 0,
            inspector: 0,
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findVideo({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.revideo
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
          res = await createVideo(this.formData)
          break
        case 'update':
          res = await updateVideo(this.formData)
          break
        default:
          res = await createVideo(this.formData)
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