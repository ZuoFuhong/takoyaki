<template>
  <el-container>
    <el-main>
      <el-row type="flex" justify="space-between">
        <el-form :inline="true" :model="searchForm" size="small">
          <el-form-item label="页面">
            <el-input v-model="searchForm.name" placeholder="请输入页面名"/>
          </el-form-item>
          <el-form-item label="数据源">
            <el-input v-model="searchForm.data_source" placeholder="请输入数据源"/>
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
        <el-button
          type="plain"
          icon="el-icon-edit-outline"
          size="small"
          @click="doJumpDataPage"
          >管理数据</el-button
        >
      </el-row>
      <!-- 表格域 -->
      <el-row class="table-field">
        <el-table border ref="singleTable" :data="tableData" highlight-current-row style="width: 100%" size="small" @current-change="handleSelectChange">
          <el-table-column type="index" width="50"></el-table-column>
          <el-table-column prop="name" label="页面名称" show-overflow-tooltip></el-table-column>
          <el-table-column prop="data_source" label="数据源" show-overflow-tooltip></el-table-column>
          <el-table-column prop="table_name" label="表名" show-overflow-tooltip></el-table-column>
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
            <el-form-item label="页面名称" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.name" autocomplete="off" :disabled="dialogFormTitle === '修改页面'"></el-input>
            </el-form-item>
            <el-form-item label="数据源" :label-width="formLabelWidth">
              <el-select v-model="dialogForm.data_source" placeholder="请选择">
                <el-option
                  v-for="item in optionDataSource"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="表名" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.table_name" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="主键" :label-width="formLabelWidth">
              <el-input v-model="dialogForm.primary_key" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="表字段" :label-width="formLabelWidth">
              <el-input type="textarea" rows="10" v-model="dialogForm.columns" :placeholder="columnsCaseHint" autocomplete="off"></el-input>
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
import pagem from '../models/page'
import ds from '../models/datasource'

export default {
  name: 'DataSource',
  data() {
    return {
      searchForm: {
        name: '',
        data_source: ''
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
      dialogForm: {},
      // 可选的数据源
      optionDataSource: []
    }
  },
  computed: {
    columnsCaseHint() {
      let casestr = JSON.stringify([
        {
          column_name: "列1",
          field_name: "c1",
          search: true
        }
      ], null, 2)
      return "表字段示例：\n" + casestr
    }
  },
  async created() {
    await this.loadPages()
    await this.loadOptionDataSource()
  },
  methods: {
    async loadOptionDataSource() {
      let {data} = await ds.getAllDataSource()
      for (var i = 0; i < data.length; i++) {
        this.optionDataSource.push(
          {
            value: data[i],
            label: data[i]
          }
        )
      }
    },
    async loadPages() {
      let form = this.searchForm
      let {retcode, errmsg, data} = await pagem.getPages({
        name: form.name,
        data_source: form.data_source,
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
      this.dialogFormTitle = "新增页面";
      this.dialogFormVisible = true;
      // 置空
      this.dialogForm = {};
    },
    async doEditRecord() {
      this.dialogFormTitle = "修改页面";
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要编辑的页面!',
          type: 'warning'
        });
        return
      }
      // deep copy
      let form = JSON.parse(JSON.stringify(this.currentRow));
      if (form["columns"]) {
        form["columns"] = JSON.stringify(form["columns"], null, 4)
      }
      this.dialogForm = form
      this.dialogFormVisible = true;
    },
    async doDeleleRecord() {
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要删除的页面!',
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
        let {retcode, errmsg} = await pagem.deletePage(this.currentRow.name)
        if (retcode !== 0) {
          this.$message.error(errmsg)
          return
        }
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
        await this.loadPages();
      }
    },
    async doDialogFormSave() {
      // deep copy
      let form = JSON.parse(JSON.stringify(this.dialogForm));
      if (!form["create_time"]) {
        form["create_time"] = this.getNowTime()
      }
      form["update_time"] = this.getNowTime()
      if (form["columns"]) {
        // 临时兼容下游
        form["columns"] = JSON.parse(form["columns"])
      }
      let {retcode, errmsg} = await pagem.savePage(form)
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      this.$message({
        message: '保存成功',
        type: 'success'
      })
      await this.loadPages();
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
      await this.loadPages();
    },
    async handleSizeChange(val) {
      // 更新单页条数
      this.pageSize = val;
      await this.loadPages();
    },
    async handleCurrentChange(val) {
      // 更新当前页码
      this.currentPage = val;
      await this.loadPages();
    },
    getNowTime() {
      return moment().format("YYYY-MM-DD hh:mm:ss")
    },
    // 跳转数据页
    doJumpDataPage() {
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要查看的页面!',
          type: 'warning'
        });
        return
      }
      let path = '/datatable?page_name=' + this.currentRow.name
      window.location.href = path
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