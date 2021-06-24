<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="内容:">
    <el-input v-model="formData.body" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="日期:">
    
      <el-date-picker type="date" placeholder="选择日期" v-model="formData.date" clearable></el-date-picker>
    </el-form-item>
    
      <el-form-item label="屏蔽:">
    <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.disabled" clearable ></el-switch>
    </el-form-item>
    
      <el-form-item label="发布者id:">
    <el-input v-model.number="formData.authorId" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="视频id:">
    <el-input v-model.number="formData.videoId" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="举报者id:">
    <el-input v-model.number="formData.tipUser" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="被举报:">
    <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.tipped" clearable ></el-switch>
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
  createComment,
  updateComment,
  findComment
} from '@/api/comment' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Comment',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            body: '',
            date: new Date(),
            disabled: false,
            authorId: 0,
            videoId: 0,
            tipUser: 0,
            tipped: false,
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findComment({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.recomment
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
          res = await createComment(this.formData)
          break
        case 'update':
          res = await updateComment(this.formData)
          break
        default:
          res = await createComment(this.formData)
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