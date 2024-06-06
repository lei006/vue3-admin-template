<template>
    <div class="layout-container">
        report_print
    </div>
</template>


<script setup>
import { ref, reactive, onMounted } from 'vue'
import apiReportPrint from '@/api/report_print'

console.log("---------------apiReportPrint:", apiReportPrint);

onMounted(()=>{
    console.log("apiReportPrint:", apiReportPrint);

})


</script>