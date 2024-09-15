<template>
    <div class="layout-container">
        <div class="layout-container-form flex space-between">
            <div class="layout-container-form-handle">
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
                <el-table-column prop="CreatedAt" label="时间" align="center" width="150" show-overflow-tooltip>
                <template #default="scope">
                      {{ convertISOToReadable(scope.row.CreatedAt) }}
                  </template>
                </el-table-column>                
                <el-table-column prop="typ" label="typ" align="center" width="100" show-overflow-tooltip />
                <el-table-column prop="msgtext01" label="msgtext01" align="center" width="240" show-overflow-tooltip />
                <el-table-column prop="msgtext02" label="msgtext02" align="center" width="240" show-overflow-tooltip />
                <el-table-column prop="msgtext03" label="msgtext03" align="center" width="240" show-overflow-tooltip />
                <el-table-column prop="msgtext04" label="msgtext04" align="center" show-overflow-tooltip />
                <el-table-column  label="来源" align="center" fixed="right" width="250" show-overflow-tooltip>
                  <template #default="scope">
                      {{ scope.row.fromip }}
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
import stringRandom  from 'string-random'
import apiSystemOption from "@/api/system/option";

const refLayout = ref()



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
  apiSystemOption.GetPage(params).then((res) => {
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

    apiSystemOption.DeleteMany(ids).then((res) => {
      ElMessage({type: "success",message: "删除成功",});
      getTableData(tableData.value.length === 1 ? true : false);
    });
}

const convertISOToReadable = (ISOString)=>{
  const date = new Date(ISOString);
  return date.Format("yyyy-MM-dd hh:mm:ss")
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
