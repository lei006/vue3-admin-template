import { defineStore } from 'pinia' //引入

const PiniaStore = defineStore('storeDictionary', {
  state: () => {
    return {
        //字段类型
        FieldType:[
          {value:"input", label:"单行输入"},
          {value:"textarea", label:"多行输入"},
          {value:"number", label:"计数器"},
          {value:"radio", label:"单选项"},
          {value:"checkbox", label:"多选项"},
          {value:"select", label:"下拉选项"},
          {value:"time", label:"时间"},
          {value:"time-range", label:"时间范围"},
          {value:"date", label:"日期"},
          {value:"date-range", label:"日期范围"},
          {value:"switch", label:"开关"},
          //{value:"color", label:"颜色"},
          {value:"rich-editor", label:"富文本"},
          {value:"checkbox", label:"多选项"},
          {value:"cascader", label:"级联选择"},
        ],
        //预置字段
        PresetFields:[],
        fromLayouts:[], //报告填写布局
        templateTree:[],
        report_setup_items:[]
    }
  },

  getters: {
    Logged(store){
      return store.token != "";
    },
  },
  actions: {

    clear(){
      this.data = null;
    },
    getFieldTypeLabel(field_type){
      for (let index = 0; index < this.FieldType.length; index++) {
        const type_element = this.FieldType[index];
        if(type_element.value == field_type) {
          return type_element.label;
        }
      }
      return "unknown";
    },
    setReportFromLayout(layouts) {
      this.fromLayouts = layouts;
    },
    setReportSetupItems(items){
      this.report_setup_items = items;
    }
  },
  //persist:true,
})
export default PiniaStore //导出
