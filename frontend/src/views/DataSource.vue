<template>
  <el-container>
    <el-main>
      <el-row type="flex" justify="space-between">
        <el-form :inline="true" :model="searchForm" size="small">
          <el-form-item label="数据源">
            <el-input v-model="searchForm.name" placeholder="请输入数据源"/>
          </el-form-item>
          <el-form-item label="主机">
            <el-input v-model="searchForm.host" placeholder="请输入主机名"/>
          </el-form-item>
          <el-form-item>
            <el-button type="plain" icon="el-icon-search" size="small" @click="onSearch">搜索</el-button>
          </el-form-item>
        </el-form>
      </el-row>
      <el-row type="flex" justify="start">
        <el-button
          type="plain"
          icon="el-icon-plus"
          size="small"
          @click="doAddRecord"
          >添加</el-button
        >
        <el-button
          type="plain"
          icon="el-icon-edit"
          size="small"
          @click="doEditRecord"
          >修改</el-button
        >
        <el-button
          type="plain"
          icon="el-icon-delete"
          size="small"
          @click="doDeleleRecord"
          >删除</el-button
        >
      </el-row>
      <!-- 表格域 -->
      <el-row class="table-field">
        <el-table border ref="singleTable" :data="tableData" highlight-current-row style="width: 100%" size="small" @current-change="handleSelectChange">
          <el-table-column type="index" width="50"></el-table-column>
          <el-table-column prop="name" label="数据源" show-overflow-tooltip></el-table-column>
          <el-table-column prop="host" label="主机" show-overflow-tooltip></el-table-column>
          <el-table-column prop="port" label="端口" show-overflow-tooltip></el-table-column>
          <el-table-column prop="database" label="数据库" show-overflow-tooltip></el-table-column>
          <el-table-column prop="create_time" label="创建时间" show-overflow-tooltip></el-table-column>
          <el-table-column prop="update_time" label="修改时间" show-overflow-tooltip></el-table-column>
        </el-table>
      </el-row>
      <!-- 分页 -->
      <el-row type="flex" justify="start" class="pagination-field">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :page-sizes="[10, 25, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="recordTotal"
        >
        </el-pagination>
      </el-row>
      <!-- 弹出层: 数据表单 -->
      <el-row type="flex" justify="start">
        <el-dialog :title="dialogFormTitle" :visible.sync="dialogFormVisible" custom-class="dialog-form">
          <el-form :model="dialogForm">
            <el-form-item label="数据源" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.name" autocomplete="off" :disabled="dialogFormTitle === '修改数据源'"></el-input>
            </el-form-item>
            <el-form-item label="主机" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.host" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="端口" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.port" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="数据库" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.database" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="账号" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.username" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.password" show-password></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer">
            <el-button type="plain" icon="el-icon-check" size="small" @click="doDialogFormSave">保存</el-button>
            <el-button type="plain" icon="el-icon-close" size="small" @click="doDialogFormCancel">取消</el-button>
          </div>
        </el-dialog>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
import moment from 'moment'
import ds from '../models/datasource'

export default {
  name: 'DataSource',
  data() {
    return {
      searchForm: {
        name: '',
        host: ''
      },
      tableData: [],
      recordTotal: 0,
      // 分页
      pageSize: 10,
      currentPage: 1,
      // 单选
      currentRow: null,
      // 数据表单
      dialogFormTitle: "数据表单",
      formLabelWidth: "80px",
      dialogFormVisible: false,
      dialogForm: {}
    }
  },
  created() {
    this.loadDataSource()
  },
  methods: {
    async loadDataSource() {
      let form = this.searchForm
      let {retcode, errmsg, data} = await ds.getDataSources({
        name: form.name,
        host: form.host,
        page: this.currentPage,
        page_size: this.pageSize
      })
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      this.tableData = data.list
      this.recordTotal = data.total
    },
    async doAddRecord() {
      this.dialogFormTitle = "新增数据源";
      this.dialogFormVisible = true;
      // 置空
      this.dialogForm = {};
    },
    async doEditRecord() {
      this.dialogFormTitle = "修改数据源";
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要编辑的项!',
          type: 'warning'
        });
        return
      }
      // deep copy
      this.dialogForm = JSON.parse(JSON.stringify(this.currentRow));
      this.dialogFormVisible = true;
    },
    async doDeleleRecord() {
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要删除的项!',
          type: 'warning'
        });
        return
      }
      let res = await this.$confirm('此操作将删除该记录, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(function() {
        // ignore
      });
      if (res == "confirm") {
        let {retcode, errmsg} = await ds.deleteDataSource(this.currentRow.name)
        if (retcode !== 0) {
          this.$message.error(errmsg)
          return
        }
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
        await this.loadDataSource();
      }
    },
    async doDialogFormSave() {
      let form = this.dialogForm
      if (!form["create_time"]) {
        form["create_time"] = this.getNowTime()
      }
      form["update_time"] = this.getNowTime()
      let {retcode, errmsg} = await ds.saveDataSource(form)
      if (retcode !== 0) {
        this.$message.error(errmsg)
        retcode
      }
      this.$message({
        message: '保存成功',
        type: 'success'
      })
      await this.loadDataSource();
      this.dialogFormVisible = false;
    },
    doDialogFormCancel() {
      this.dialogFormVisible = false;
      this.dialogForm = {};
    },
    handleSelectChange(val) {
      this.currentRow = val;
    },
    async onSearch() {
      this.currentPage = 1
      await this.loadDataSource();
    },
    async handleSizeChange(val) {
      // 更新单页条数
      this.pageSize = val;
      await this.loadDataSource();
    },
    async handleCurrentChange(val) {
      // 更新当前页码
      this.currentPage = val;
      await this.loadDataSource();
    },
    getNowTime() {
      return moment().format("YYYY-MM-DD hh:mm:ss")
    }
  }
}
</script>

<style lang="scss" scoped>
.el-container {
  padding: 20px;
}
.table-field {
  margin-top: 18px;
}
.pagination-field {
  margin-top: 18px;
}
.dialog-form {
  height: 70vh;
  overflow: auto;
}
</style>