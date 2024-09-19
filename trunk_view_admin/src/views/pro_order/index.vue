<template>
  <div class="layout-container">
    <div class="layout-container-form flex space-between">
      <div class="layout-container-form-handle">
        <el-select v-model="project_id" placeholder="Select" size="small" style="width: 120px; margin-right: 10px;" @change="handleProjectChange">
            <el-option v-for="item in project_items" :key="item.value" :label="item.id +' '+ item.label" :value="item.value"/>
        </el-select>
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
        <el-table-column prop="project_id" label="项目" align="center" width="160">
          <template #default="scope">
                <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'title')"> {{ getProjectName(scope.row.project_id)}} </div>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题:" align="center" width="160">
          <template #default="scope">
                <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'title')"> {{ scope.row.title?scope.row.title:"-" }} </div>
          </template>
        </el-table-column>
        
        <el-table-column prop="price" label="价格" align="center" width="160" >
          <template #default="scope">
            <el-input-number v-model="scope.row.price" @change="handleUpdateStatus(scope.row, 'price', scope.row.price)" />
          </template>
        </el-table-column>


        <el-table-column prop="status" label="状态" align="center" width="100">
          <template #default="scope">
            <el-checkbox v-model="scope.row.is_disable" label="禁用" @change="handleUpdateStatus(scope.row, 'is_disable', scope.row.is_disable)" />
          </template>
        </el-table-column>

        <el-table-column prop="remark" label="desc" align="center">
          <template #default="scope">
                <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'remark')"> {{ scope.row.remark?scope.row.remark:"-" }} </div>
          </template>          
        </el-table-column>
        <el-table-column  :label="$t('message.common.handle')" align="center" fixed="right" width="120">
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
        <el-form :model="addForm" :rules="addUserRules" ref="form" label-width="120px" style="margin-right:30px;">

            <el-form-item label="标题" prop="company">
                <el-input v-model="addForm.title" placeholder="请输入名称"></el-input>
            </el-form-item>

            <el-form-item label="禁用：" prop="company">
              <el-checkbox v-model="addForm.is_disable" label="停用" />
            </el-form-item>

            <el-form-item label="描述：" >
                <el-input v-model="addForm.remark" placeholder="请输入描述"></el-input>
            </el-form-item>
        </el-form>

        <template #footer>
            <div class="dialog-footer">
                <el-button @click="addUserDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="onBtnAdd">确定</el-button>
            </div>
        </template>
    </el-dialog>






</template>

<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";
import { Page } from "@/components/table/type";
import { ElMessage, ElMessageBox } from "element-plus";
import Table from "@/components/table/index.vue";
import { Plus, Delete, Search } from '@element-plus/icons'
import stringRandom  from 'string-random'
import apiProOrderOption from "@/api/pro/order";
import apiProProjectOption from "@/api/pro/project";

const refLayout = ref()


const project_id = ref()

const project_items = ref([])



// 存储搜索用的数据
const query = reactive({
  keyword: "",
});


const addUserDialogVisible = ref(false)
let addForm = reactive({
  project_id:0,
  title:"",
  remark: '',
  price:1,
  is_disable: false,
})

const addUserRules = {
  company: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
}

const onBtnAdd = () => {

}




// 分页参数, 供table使用
const page: Page = reactive({
  index: 1,
  size: 20,
  total: 0,
});
const loading = ref(true);
const tableData = ref([]);
const chooseData = ref([]);
const handleSelectionChange = (val: []) => {
  chooseData.value = val;
};


const getProjectName = (id)=>{

  let name = ""
  project_items.value.forEach((item: any) => {
    if (item.id == id) {
      name = "[" + id + "] " +item.label
    }
  })



  return name
}



// 获取表格数据
// params <init> Boolean ，默认为false，用于判断是否需要初始化分页
const getTableData = (init: Boolean) => {
  loading.value = true
  if (init) {
    page.index = 1
  }
  let params = {
    page: page.index,
    pageSize: page.size,
    ...query
  }
  apiProOrderOption.GetPage({project_id: project_id.value}).then((res) => {
      let data = res.data.items
      data.forEach((d: any) => {
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
const handleDel = (data: object[]) => {
    let ids = []
    data.map((e: any) => {
      ids.push(e.id);
    })

    apiProOrderOption.DeleteMany(ids).then((res) => {
      ElMessage({type: "success",message: "删除成功",});
      getTableData(tableData.value.length === 1 ? true : false);
    });
}


const handleProjectChange = ()=>{
  getTableData(false)
}


// 新增弹窗功能
const handleAdd = () => {

    addUserDialogVisible.value = true;
}


// 状态编辑功能
const handleUpdateStatus = (row, field, data) => {
    
    if (!row.id) {
      return
    }
    row.loading = true
    apiProOrderOption.PatchOne(row.id,field, data ). then(res => {
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


const handleFieldEdit = (row: object, field :string) => {

    ElMessageBox.prompt('', '修改:' + field, {inputValue: row[field], confirmButtonText: '确定',cancelButtonText: '取消'}).then(({ value }) => {
      apiProOrderOption.PatchOne(row.id, field, value).then(res => {
            if (res.code === 200) {
                row[field] = res.data[field];
            } else {
                ElMessage({type: 'error',message: '操作失败'})
            }
      })
  }).catch(() => {
      ElMessage({type: 'info',message: 'Input canceled'})
  })

}



const RefreshProjectList = ()=>{
  apiProProjectOption.GetPage().then((res) => {
      let data = res.data.items
      project_items.value = []
      data.forEach((d: any) => {
        project_items.value.push({id:d.id, label: d.project_name, value: d.id, remark: d.remark})
      })
    }).catch((error) => {
      console.log(error)
    }).finally(() => {
      
    });
}




onMounted(()=>{
    getTableData(true)
    RefreshProjectList()
})





</script>

<style lang="scss" scoped>
.statusName {
  margin-right: 10px;
}
</style>
