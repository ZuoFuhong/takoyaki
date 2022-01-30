<template>
  <div id="app">
    <el-container>
      <el-main>
        <!-- 搜索域 -->
        <el-row type="flex" justify="space-between">
          <el-form :inline="true" :model="query" size="small">
            <el-form-item v-for="field in searchField" :key="field._id" :label="field.fieldName">
              <el-input v-model="query[field.column]" :placeholder="field.fieldName"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="plain" icon="el-icon-search" size="small" @click="onSearch">搜索</el-button>
            </el-form-item>
          </el-form>
        </el-row>
        <!-- 按钮域 -->
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
            icon="el-icon-upload2"
            size="small"
            @click="doExportRecords"
            >导出</el-button
          >
        </el-row>
        <!-- 表格域 -->
        <el-row class="table-field">
          <el-table border ref="singleTable" :data="tableData" highlight-current-row style="width: 100%" size="small" @current-change="handleSelectChange">
            <el-table-column type="index" width="50"></el-table-column>
            <el-table-column v-for="field in tableField" :key="field._id" :prop="field.column" :label="field.fieldName" show-overflow-tooltip></el-table-column> 
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
              <el-form-item v-for="field in tableField" :key="field._id" :label="field.fieldName" :label-width="formLabelWidth">
                <el-input v-model="dialogForm[field.column]" autocomplete="off"></el-input>
              </el-form-item>
            </el-form>
            <div slot="footer">
              <el-button type="plain" icon="el-icon-check" size="small" @click="doDialogFormSave">保存</el-button>
              <el-button type="plain" icon="el-icon-close" size="small" @click="doDialogFormCancel">退出</el-button>
            </div>
          </el-dialog>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import rest from "@/models/rest";
import fileSaver from 'file-saver';
import xlsx from 'xlsx';

export default {
  name: "app",
  data() {
    return {
      // 搜索字段
      searchField: [],
      // 表格布局
      tableField: [],
      // 查询
      query: {
        offset: 0,
        limit: 0,
      },
      // 数据
      tableData: [],
      recordTotal: 0,   
      // 分页
      currentPage: 1,
      pageSize: 10,
      // 单选
      currentRow: null,
      // 数据表单
      dialogFormTitle: "数据表单",
      formLabelWidth: "120px",
      dialogFormVisible: false,
      dialogForm: {}
    };
  },
  async created() {
    this.initLayout();
    await this.loadRecords();
  },
  methods: {
    // 表格布局
    async initLayout() {
      let res = await rest.columns();
      let fields = res.data;
      let seq = 0;
      for (var i = 0; i < fields.length; i++) {
        let item  = fields[i];
        let field = {
          "_id": ++seq,
          "column": item.column,
          "fieldName": item.fieldName
        }
        if (item["search"]) {
          this.searchField.push(field)
        }
        this.tableField.push(field);
      }
    },
    async loadRecords() {
      // 计算分页偏移量
      this.query.offset = (this.currentPage - 1) * this.pageSize;
      this.query.limit = this.pageSize;
      let res = await rest.select(this.query);
      if (res.status != 200) {
        this.$message.error('请检查网络，数据加载失败!');
        return
      }
      this.tableData = res.data.list;
      this.recordTotal = res.data.total;
    },
    async onSearch() {
      await this.loadRecords();
    },
    async handleSizeChange(val) {
      // 更新单页条数
      this.pageSize = val;
      await this.loadRecords();
    },
    async handleCurrentChange(val) {
      // 更新当前页码
      this.currentPage = val;
      await this.loadRecords();
    },
    handleSelectChange(val) {
      this.currentRow = val;
    },
    doAddRecord() {
      this.dialogFormTitle = "新增数据";
      this.dialogFormVisible = true;
      // 置空
      this.dialogForm = {};
    },
    doEditRecord() {
      this.dialogFormTitle = "修改数据表单";
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要编译的项!',
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
        let record = JSON.parse(JSON.stringify(this.currentRow))
        let res = await rest.delete(record)
        if (res.status == 200) {
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
        } else {
          this.$message.error('删除失败');
        }
        await this.loadRecords();
      }
    },
    doExportRecords() {
      // 导出当前页
      let ws = xlsx.utils.json_to_sheet(this.tableData)
      let wb = xlsx.utils.book_new()
      xlsx.utils.book_append_sheet(wb, ws, 'sheet1')
      let wb_out = xlsx.write(wb, { bookType: 'xlsx', type: 'array' });
      try {
        fileSaver.saveAs(
          new Blob([wb_out], {
            type: 'application/octet-stream'
          }),
          'sheet.xlsx'
        )
      } catch (e) {
        console.log(e)
        this.$message.error('导出失败，请按下 F12 查看控制台日志');
      }
    },
    async doDialogFormSave() {
      let res = {};
      if (this.dialogFormTitle == "新增数据") {
        res = await rest.insert(this.dialogForm);
      } else {
        res = await rest.update(this.dialogForm);
      }
      if (res.status == 200) {
        this.$message({
          message: '保存成功',
          type: 'success'
        });
      } else {
        console.log(res)
        this.$message.error('保存失败');
      }
      await this.loadRecords();
      this.dialogFormVisible = false;
    },
    doDialogFormCancel() {
      this.dialogFormVisible = false;
      this.dialogForm = {};
    }
  },
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
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
