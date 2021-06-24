<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="用户名:">
        <el-input v-model="formData.username" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="密码:">
        <el-input v-model="formData.password" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="邮箱:">
        <el-input v-model="formData.email" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="头像:">
        <el-input v-model="formData.avatar" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="昵称:">
        <el-input v-model="formData.nickname" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="简介:">
        <el-input v-model="formData.aboutMe" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="注册时间:">

        <el-date-picker v-model="formData.memberSince" type="date" placeholder="选择日期" clearable />
      </el-form-item>

      <el-form-item label="最近登录:">

        <el-date-picker v-model="formData.lastSeen" type="date" placeholder="选择日期" clearable />
      </el-form-item>

      <el-form-item label="权限:">
        <el-input v-model.number="formData.permissions" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="硬币:">
        <el-input v-model.number="formData.coins" clearable placeholder="请输入" />
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
  createUsers,
  updateUsers,
  findUsers
} from '@/api/users' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Users',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        username: '',
        password: '',
        email: '',
        avatar: '',
        nickname: '',
        aboutMe: '',
        memberSince: new Date(),
        lastSeen: new Date(),
        permissions: 0,
        coins: 0

      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findUsers({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reusers
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
          res = await createUsers(this.formData)
          break
        case 'update':
          res = await updateUsers(this.formData)
          break
        default:
          res = await createUsers(this.formData)
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
