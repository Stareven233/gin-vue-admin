<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="openDialog">新增评论</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="40" />
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="内容" prop="body" width="320" />
      <el-table-column label="发布日期" prop="date" width="120">
        <template slot-scope="scope">{{ scope.row.date|formatDate2 }}</template>
      </el-table-column>
      <el-table-column label="屏蔽" prop="disabled" width="80">
        <template slot-scope="scope">{{ scope.row.disabled|formatBoolean }}</template>
      </el-table-column>
      <el-table-column label="发布者" prop="authorId" width="80" />
      <el-table-column label="视频" prop="videoId" width="80" />
      <el-table-column label="举报者" prop="tipUser" width="80" />
      <el-table-column label="被举报" prop="tipped" width="80">
        <template slot-scope="scope">{{ scope.row.tipped|formatBoolean }}</template>
      </el-table-column><el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateComment(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="内容:">

          <el-input v-model="formData.body" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发布日期:">

          <el-date-picker v-model="formData.date" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="屏蔽:">

          <el-switch v-model="formData.disabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="发布者">

          <el-input v-model.number="formData.authorId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="视频">

          <el-input v-model.number="formData.videoId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报者">

          <el-input v-model.number="formData.tipUser" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="被举报:">

          <el-switch v-model="formData.tipped" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createComment,
  deleteComment,
  deleteCommentByIds,
  updateComment,
  findComment,
  getCommentList
} from '@/api/comment' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'Comment',
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    formatDate2: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd')
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  mixins: [infoList],
  data() {
    return {
      listApi: getCommentList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],

      formData: {
        body: '',
        date: new Date(),
        disabled: false,
        authorId: 0,
        videoId: 0,
        tipUser: 0,
        tipped: false

      }
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.disabled === '') {
        this.searchInfo.disabled = null
      }
      if (this.searchInfo.tipped === '') {
        this.searchInfo.tipped = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteComment(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteCommentByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateComment(row) {
      const res = await findComment({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recomment
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        body: '',
        date: new Date(),
        disabled: false,
        authorId: 0,
        videoId: 0,
        tipUser: 0,
        tipped: false

      }
    },
    async deleteComment(row) {
      const res = await deleteComment({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
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
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  }
}
</script>

<style lang="scss" scoped>
  .el-table .el-button {
    margin: 3px 4px;
  }
</style>
