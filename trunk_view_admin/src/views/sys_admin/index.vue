<template>
  <div class="layout-container">
    <div class="layout-container-form flex space-between">
      <div class="layout-container-form-handle">
        <el-button type="primary" :icon="Plus" @click="handleAdd">{{
          $t("message.common.add")
        }}</el-button>
        <el-popconfirm
          :title="$t('message.common.delTip')"
          @confirm="handleDel(chooseData)"
        >
          <template #reference>
            <el-button
              type="danger"
              :icon="Delete"
              :disabled="chooseData.length === 0"
              >{{ $t("message.common.delBat") }}</el-button
            >
          </template>
        </el-popconfirm>
      </div>
      <div class="layout-container-form-search">
        <el-input
          v-model="query.keyword"
          :placeholder="$t('message.common.searchTip')"
        ></el-input>
        <el-button
          type="primary"
          :icon="Search"
          class="search-btn"
          @click="getTableData(true)"
          >{{ $t("message.common.search") }}</el-button
        >
      </div>
    </div>
    <div class="layout-container-table">
      <Table
        ref="table"
        v-model:page="page"
        v-loading="loading"
        :showSelection="true"
        :data="tableData"
        @getTableData="getTableData"
        @selection-change="handleSelectionChange"
      >
        <el-table-column prop="id" label="Id" align="center" width="60" />
        <el-table-column prop="username" label="用户名" align="center" width="160"/>
        <el-table-column prop="nickname" label="昵称" align="center"  width="120">
            <template #default="scope">
                <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'nickname')"> {{ scope.row.nickname?scope.row.nickname:"-" }} </div>
            </template>
        </el-table-column>
        <el-table-column prop="password" label="密码" align="center"  width="120">
            <template #default="scope">
                <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'password')"> {{ "-" }} </div>
            </template>
        </el-table-column>

        <el-table-column prop="usersign" label="签名" align="center"  width="180">
            <template #default="scope">
                <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'usersign')"> {{ scope.row.usersign?scope.row.usersign:"-" }} </div>
            </template>            
        </el-table-column>
        <el-table-column prop="status" label="状态" align="center" width="100">
          <template #default="scope">
            <el-checkbox v-model="scope.row.is_disable" label="禁用" @change="handleUpdateStatus(scope.row, 'is_disable', scope.row.is_disable)" />
          </template>
        </el-table-column>
        <el-table-column prop="desc" label="desc" align="center">
            <template #default="scope">
                <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'desc')"> {{ scope.row.desc?scope.row.desc:"-" }} </div>
            </template>
        </el-table-column>
        <el-table-column  :label="$t('message.common.handle')" align="center" fixed="right" width="100">
          <template #default="scope">
            <el-popconfirm :title="$t('message.common.delTip')" @confirm="handleDel([scope.row])">
              <template #reference>
                <el-button type="danger">{{ $t("message.common.del") }}</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </Table>
    </div>
  </div>


    <el-dialog v-model="addUserDialogVisible" title="Tips" width="500">
        <el-form :model="addUserForm" :rules="addUserRules" ref="form" label-width="120px" style="margin-right:30px;">
            <el-form-item label="名称：" prop="username">
                <el-input v-model="addUserForm.username" placeholder="请输入名称"></el-input>
            </el-form-item>
            <el-form-item label="昵称：" >
                <el-input v-model="addUserForm.nickname" placeholder="请输入昵称"></el-input>
            </el-form-item>
            <el-form-item label="密码：" prop="password">
                <el-input v-model="addUserForm.password" placeholder="请输入密码"></el-input>
            </el-form-item>
            <el-form-item label="重录密码" prop="password1">
                <el-input v-model="addUserForm.password1" placeholder="请再次输入密码"></el-input>
            </el-form-item>

            <el-form-item label="签名：" >
                <el-input v-model="addUserForm.usersign" placeholder="请输入签名"></el-input>
            </el-form-item>

            <el-form-item label="是否停用:" >
                <el-checkbox v-model="addUserForm.is_disable" label="停用" />
            </el-form-item>

            <el-form-item label="描述：" >
                <el-input v-model="addUserForm.desc" placeholder="请输入描述"></el-input>
            </el-form-item>
        </el-form>

        <template #footer>
            <div class="dialog-footer">
                <el-button @click="addUserDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="onBtnAddUser">确定</el-button>
            </div>
        </template>
    </el-dialog>






</template>

<script setup>
import { onMounted, ref, reactive } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import Table from "@/components/table/index.vue";
import { Plus, Delete, Search } from '@element-plus/icons'
import stringRandom  from 'string-random'
import apiSystemAdmin from "@/api/system/admin";

const refLayout = ref()

// 存储搜索用的数据
const query = reactive({
  keyword: "",
});




const addUserDialogVisible = ref(false)
let addUserForm = reactive({
  username: '',
  nickname: '',
  password: '',
  password1: '',
  usersign: '',
  is_disable: false,
  desc: '',
})

const addUserRules = {
  username: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  password1: [{ required: true, message: '确认密码', trigger: 'blur' }],
  sort: [{ required: true, message: '请输入数字', trigger: 'blur' }],
  select: [{ required: true, message: '请选择', trigger: 'blur' }],
  radio: [{ required: true, message: '请选择', trigger: 'blur' }]
}

const onBtnAddUser = () => {

  apiSystemAdmin.AddOne(addUserForm).then(res => {
        // 把 res.data 插入到tableData的前面
        tableData.value.unshift(res.data);
        ElMessage({type:'success',message: '新增成功'})
        addUserDialogVisible.value = false;
    })
}




// 分页参数, 供table使用
const page = reactive({
  index: 1,
  size: 20,
  total: 0,
});
const loading = ref(true);
const tableData = ref([]);
const chooseData = ref([]);
const handleSelectionChange = (val) => {
  chooseData.value = val;
};





// 获取表格数据
// params <init> Boolean ，默认为false，用于判断是否需要初始化分页
const getTableData = (init) => {
  loading.value = true
  if (init) {
    page.index = 1
  }
  let params = {
    page: page.index,
    pageSize: page.size,
    ...query
  }
  apiSystemAdmin.GetPage(params).then((res) => {
      let data = res.data.items
      data.forEach((d) => {
        d.loading = false
      })
      tableData.value = data
      page.total = Number(res.data.total);
    }).catch((error) => {
      tableData.value = [];
      page.index = 1;
      page.total = 0;
    }).finally(() => {
      loading.value = false;
    });
}
  // 删除功能
const handleDel = (data) => {
    let ids = []
    data.map((e) => {
      ids.push(e.id);
    })

    apiSystemAdmin.DeleteMany(ids).then((res) => {
      ElMessage({type: "success",message: "删除成功",});
      getTableData(tableData.value.length === 1 ? true : false);
    });
}


// 新增弹窗功能
const handleAdd = () => {


    addUserForm.username = stringRandom()
    addUserForm.nickname = stringRandom()
    addUserForm.password = stringRandom()
    addUserForm.password1 = stringRandom()
    addUserForm.usersign = stringRandom()


    addUserDialogVisible.value = true;

    //console.log("layer", refLayout)
    //refLayout.value.showAdd();

    //getTableData(tableData.value.length === 1 ? true : false);

    //layer.title = "新增数据";
    //layer.show = true;
    //delete layer.row;
}
// 编辑弹窗功能
const handleEdit = (row) => {

    refLayout.value.showEdit(row);
  
  
}
// 状态编辑功能
const handleUpdateStatus = (row, field, data) => {
    
    if (!row.id) {
      return
    }
    row.loading = true
    apiSystemAdmin.PatchOne(row.id,field, data ). then(res => {
        if (res.code === 200) {
            row[field] = res.data[field];
        } else {
            ElMessage({type: 'error',message: '操作失败'})
        }
    }).catch(err => {
      ElMessage({type: 'error',message: '操作失败'})
    }).finally(() => {
      row.loading = false
    })
    


    
}


const handleFieldEdit = (row, field) => {

    ElMessageBox.prompt('', '修改:' + field, {inputValue: row[field], confirmButtonText: '确定',cancelButtonText: '取消'}).then(({ value }) => {
      apiSystemAdmin.PatchOne(row.id, field, value).then(res => {
            if (res.code === 200) {
                row[field] = res.data[field];
            } else {
                ElMessage({type: 'error',message: '操作失败'})
            }
      })
  })
  .catch(() => {
      ElMessage({type: 'info',message: 'Input canceled'})
  })

}








onMounted(()=>{
    getTableData(true)
})





</script>

<style lang="scss" scoped>
.statusName {
  margin-right: 10px;
}
</style>
