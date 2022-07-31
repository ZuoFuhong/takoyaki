<template>
  <el-container>
    <el-main>
      <!-- 搜索域 -->
      <el-row type="flex" justify="space-between">
        <el-form :inline="true" :model="queryForm" size="small">
          <el-form-item v-for="field in searchField" :key="field._id" :label="field.column_name">
            <el-input v-model="queryForm[field.field_name]" :placeholder="field.column_name"></el-input>
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
          <el-table-column v-for="field in tableField" :key="field._id" :prop="field.field_name" :label="field.column_name" show-overflow-tooltip></el-table-column> 
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
            <el-form-item v-for="field in tableField" :key="field._id" :label="field.column_name" :label-width="formLabelWidth">
              <el-input v-model="dialogForm[field.field_name]" autocomplete="off"></el-input>
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
import rest from "@/models/rest";
import fileSaver from 'file-saver';
import xlsx from 'xlsx';

export default {
  name: "DataTable",
  data() {
    return {
      pageName: "书籍页",
      // 搜索字段
      searchField: [],
      // 表格布局
      tableField: [],
      // 查询条件
      queryForm: {},
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
    let pageName = this.$route.query.page_name;
    if (pageName) {
      this.pageName = pageName
      this.initLayout();
      await this.loadRecords();
    }
  },
  methods: {
    // 表格布局
    async initLayout() {
      let {retcode, errmsg, data} = await rest.columns(this.pageName);
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      let seq = 0;
      for (var i = 0; i < data.length; i++) {
        let item  = data[i];
        let field = {
          "_id": ++seq,
          "column_name": item.column_name,
          "field_name": item.field_name
        }
        if (item["search"]) {
          this.searchField.push(field)
        }
        this.tableField.push(field);
      }
    },
    async loadRecords() {
      let {retcode, errmsg, data} = await rest.select({
        conditions: this.queryForm,
        page_name: this.pageName,
        page: this.currentPage,
        page_size: this.pageSize
      })
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      this.tableData = data.list;
      this.recordTotal = data.total;
    },
    async onSearch() {
      this.currentPage = 1
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
        let {retcode, errmsg} = await rest.delete({
          page_name: this.pageName,
          record: this.currentRow
        })
        if (retcode !== 0) {
          this.$message.error(errmsg)
          return
        }
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
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
      if (this.dialogFormTitle == "新增数据") {
        let {retcode, errmsg} = await rest.insert({
          page_name: this.pageName,
          record: this.dialogForm
        })
        if (retcode !== 0) {
          this.$message.error(errmsg)
          return
        }
      } else {
        let {retcode, errmsg} = await rest.update({
          page_name: this.pageName,
          record: this.dialogForm
        })
        if (retcode !== 0) {
          this.$message.error(errmsg)
          return
        }
      }
      this.$message({
        message: '保存成功',
        type: 'success'
      });
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
