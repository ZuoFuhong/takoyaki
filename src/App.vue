<template>
  <div id="app">
    <el-container>
      <el-main>
        <!-- 搜索域 -->
        <el-row type="flex" justify="space-between">
          <el-form :inline="true" :model="query" size="small">
            <el-form-item label="书名">
              <el-input v-model="query.bookName" placeholder="书名"></el-input>
            </el-form-item>
            <el-form-item label="作者">
              <el-input v-model="query.author" placeholder="作者"></el-input>
            </el-form-item>
            <el-form-item label="状态">
              <el-select v-model="query.status" placeholder="状态">
                <el-option label="上架" value="1"></el-option>
                <el-option label="下架" value="0"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button
                type="plain"
                icon="el-icon-search"
                size="small"
                @click="onSearch"
                >搜索</el-button
              >
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
          <el-table
            border
            ref="singleTable"
            :data="tableData"
            highlight-current-row
            style="width: 100%"
            size="small"
            @current-change="handleSelectChange"
          >
            <el-table-column type="index" width="50"></el-table-column>
            <el-table-column prop="isbn" label="ISBN"></el-table-column>
            <el-table-column prop="name" label="书名"></el-table-column>
            <el-table-column prop="price" label="定价"></el-table-column>
            <el-table-column prop="author" label="作者"></el-table-column>
            <el-table-column prop="edition" label="版次"></el-table-column>
            <el-table-column prop="press" label="出版社"></el-table-column>
            <el-table-column
              prop="address"
              label="社址"
              show-overflow-tooltip
            ></el-table-column>
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
              <el-form-item label="ISBN" :label-width="formLabelWidth">
                <el-input v-model="dialogForm.isbn" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="书名" :label-width="formLabelWidth">
                <el-input v-model="dialogForm.name" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="定价" :label-width="formLabelWidth">
                <el-input v-model="dialogForm.price" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="作者" :label-width="formLabelWidth">
                <el-input v-model="dialogForm.author" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="版次" :label-width="formLabelWidth">
                <el-input v-model="dialogForm.edition" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="出版社" :label-width="formLabelWidth">
                <el-input v-model="dialogForm.press" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="社址" :label-width="formLabelWidth">
                <el-input v-model="dialogForm.address" autocomplete="off"></el-input>
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
import storage from "@/models/storage";
import fileSaver from 'file-saver';
import xlsx from 'xlsx';

export default {
  name: "app",
  data() {
    return {
      query: {
        bookName: "",
        author: "",
        status: "",
        offset: 0,
        limit: 0,
      },
      tableData: [],      
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
  computed: {
    recordTotal() {
      return this.tableData.length
    }
  },
  async created() {
    await this.loadRecords();
  },
  methods: {
    async loadRecords() {
      // 计算分页偏移量
      this.query.offset = (this.currentPage - 1) * this.pageSize;
      this.query.limit = this.pageSize;
      let records = await storage.select(this.query);
      this.tableData = records;
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
      this.dialogFormVisible = true;
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要编译的项!',
          type: 'warning'
        });
        return
      }
      // deep copy
      this.dialogForm = JSON.parse(JSON.stringify(this.currentRow));
    },
    doDeleleRecord() {
      if (this.currentRow == null) {
        this.$message({
          message: '请先选择要删除的项!',
          type: 'warning'
        });
        return
      }
      this.$confirm('此操作将删除该记录, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        let record = JSON.parse(JSON.stringify(this.currentRow))
        storage.delete(record)
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
      }).catch(() => {
        // ignore catch
      });
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
    doDialogFormSave() {
      if (this.dialogFormTitle == "新增数据") {
        storage.insert(this.dialogForm);
      } else {
        storage.update(this.dialogForm);
      }
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
