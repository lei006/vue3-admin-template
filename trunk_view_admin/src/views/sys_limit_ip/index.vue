<template>
    <div class="layout-container">
        <div class="layout-container-form flex space-between">
            <div class="layout-container-form-handle">

                <el-button type="primary" :icon="Plus" @click="handleAdd">{{$t("message.common.add")}}</el-button>              
                <el-popconfirm :title="$t('message.common.delTip')" @confirm="handleDel(chooseData)">
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
                <el-table-column prop="id" label="id" align="center" width="100" show-overflow-tooltip />
                <el-table-column prop="CreatedAt" label="创建时间" align="center" width="240" show-overflow-tooltip />
                <el-table-column prop="ip" label="要限制的IP" align="center" width="240" show-overflow-tooltip />
                <el-table-column prop="status" label="是否限制" align="center" width="100">
                  <template #default="scope">
                    <el-checkbox v-model="scope.row.is_limit" label="限制" @change="handleUpdateStatus(scope.row, 'is_limit', scope.row.is_limit)" />
                  </template>
                </el-table-column>

                <el-table-column prop="desc" label="描述" align="center" show-overflow-tooltip>
                  <template #default="scope">
                    <div style="line-height: 32px; cursor: pointer; background-color: rgba(100, 100, 100, 0.05);" @click="handleFieldEdit(scope.row, 'desc')"> {{ scope.row.desc?scope.row.desc:"-" }} </div>
                  </template>
                </el-table-column>
            </Table>
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref, reactive } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import Table from "@/components/table/index.vue";
import { Plus, Delete, Search } from '@element-plus/icons'
import apiSystemLimitIp from "@/api/system/limit_ip";


// 存储搜索用的数据
const query = reactive({
  keyword: "",
});





// 分页参数, 供table使用
const page = reactive({
  index: 1,
  size: 100,
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
  apiSystemLimitIp.GetPage(params).then((res) => {
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

    apiSystemLimitIp.DeleteMany(ids).then((res) => {
      ElMessage({type: "success",message: "删除成功",});
      getTableData(tableData.value.length === 1 ? true : false);
    });
}


// 新增弹窗功能
const handleAdd = () => {

  ElMessageBox.prompt('', '修改IP:', {inputValue: "", confirmButtonText: '确定',cancelButtonText: '取消'}).then(({ value }) => {
      apiSystemLimitIp.AddOne({ip: value}).then(res => {
            if (res.code === 200) {
              tableData.value.unshift(res.data);                
            } else {
                ElMessage({type: 'error',message: '操作失败'})
            }
      })
  }).catch(() => {
      ElMessage({type: 'info',message: 'Input canceled'})
  })
}


const handleFieldEdit = (row, field) => {

    ElMessageBox.prompt('', '修改:' + field, {inputValue: row[field], confirmButtonText: '确定',cancelButtonText: '取消'}).then(({ value }) => {
        apiSystemLimitIp.PatchOne(row.id, field, value).then(res => {
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


const handleUpdateStatus = (row, field, data) => {
    
    if (!row.id) {
      return
    }

    row.loading = true
    apiSystemLimitIp.PatchOne(row.id,field, data ). then(res => {
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





onMounted(()=>{
    getTableData(true)
})





</script>

<style lang="scss" scoped>
.statusName {
  margin-right: 10px;
}
</style>
