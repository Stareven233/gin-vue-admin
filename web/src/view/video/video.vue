<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="title字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.title" />
        </el-form-item>                          
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="openDialog">新增视频</el-button>
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
      <el-table-column type="selection" width="55" />
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="title字段" prop="title" width="120" /> 
      <el-table-column label="file字段" prop="file" width="120" /> 
      <el-table-column label="date字段" prop="date" width="120" /> 
      <el-table-column label="authorId字段" prop="authorId" width="120" /> 
      <el-table-column label="type字段" prop="type" width="120" /> 
      <el-table-column label="face字段" prop="face" width="120" /> 
      <el-table-column label="desc字段" prop="desc" width="120" /> 
      <el-table-column label="coin字段" prop="coin" width="120" /> 
      <el-table-column label="collect字段" prop="collect" width="120" /> 
      <el-table-column label="like字段" prop="like" width="120" /> 
      <el-table-column label="status字段" prop="status" width="120" /> 
      <el-table-column label="inspector字段" prop="inspector" width="120" /> <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateVideo(scope.row)">变更</el-button>
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
        <el-form-item label="title字段:">
      
          <el-input v-model="formData.title" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="file字段:">
      
          <el-input v-model="formData.file" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="date字段:">
      
          <el-date-picker type="date" placeholder="选择日期" v-model="formData.date" clearable />
       </el-form-item>
        <el-form-item label="authorId字段:">
      
          <el-input v-model.number="formData.authorId" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="type字段:">
      
          <el-input v-model.number="formData.type" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="face字段:">
      
          <el-input v-model="formData.face" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="desc字段:">
      
          <el-input v-model="formData.desc" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="coin字段:">
      
          <el-input v-model.number="formData.coin" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="collect字段:">
      
          <el-input v-model.number="formData.collect" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="like字段:">
      
          <el-input v-model.number="formData.like" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="status字段:">
      
          <el-input v-model.number="formData.status" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="inspector字段:">
      
          <el-input v-model.number="formData.inspector" clearable placeholder="请输入" />
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
  createVideo,
  deleteVideo,
  deleteVideoByIds,
  updateVideo,
  findVideo,
  getVideoList
} from '@/api/video' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'Video',
  mixins: [infoList],
  data() {
    return {
      listApi: getVideoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
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
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time);
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss');
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
  async created() {
    await this.getTableData()
    
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10            
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
        this.deleteVideo(row)
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
      const res = await deleteVideoByIds({ ids })
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
    async updateVideo(row) {
      const res = await findVideo({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.revideo
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
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
    },
    async deleteVideo(row) {
      const res = await deleteVideo({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1 ) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case "create":
          res = await createVideo(this.formData)
          break
        case "update":
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
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
