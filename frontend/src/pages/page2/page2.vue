<template>
<div class="title" >Протокол
      <v-data-table
      class=" blue-grey lighten-5 table table-bordered table-hover "  border="2" height="10%" with="100%"
      dense
      :headers="headers"
      :items="protocol"
      :items-per-page="30"
      :rows-per-page-items='[30,10,{text: "Все", value: -1}]'
      >
      </v-data-table>
</div>
</template>
<script>
let colorHeader = 'blue lighten-4'
export default {
  data: () => ({
    headers: [
      {
        text: 'Дата',
        sortable: false,
        value: 'VcDate',
        class: colorHeader
      },
      { text: 'Начало', value: 'CBegTime', class: colorHeader, sortable: false },
      { text: 'Конец', value: 'CEndTime', class: colorHeader, sortable: false },
      { text: 'Всего', value: 'CDeltaTime', class: colorHeader, sortable: false },
      { text: 'Забой', value: 'FEndDrillDepth', class: colorHeader, sortable: false },
      { text: 'Пол. долота', value: 'FEndToolDepth', class: colorHeader, sortable: false },
      { text: 'Код', value: 'CZatrName', class: colorHeader, sortable: false },
      { text: 'Операция', value: 'VcGroupOper', class: colorHeader, sortable: false },
      { text: 'Описание работ', value: 'VcComment', class: colorHeader, width: 200 },
      { text: 'Рейс', value: 'IRangeNum', class: colorHeader, sortable: false },
      { text: 'Этап', value: 'IStageNum', class: colorHeader, sortable: false },
      { text: 'Ок', value: 'ICasingNum', class: colorHeader, sortable: false },
      { text: 'Парам. р-ра', value: 'VcSolutions', class: colorHeader, sortable: false }
    ],
    protocol: [
      {
        VcDate: '11.11.19',
        CBegTime: '00:00',
        CEndTime: '01:00',
        CDeltaTime: '01:00',
        FEndDrillDepth: 1000,
        FEndToolDepth: 1000,
        CZatrName: 'бурение',
        VcGroupOper: 'бурение',
        VcComment: 'Бурение в инт. 900-1100м',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '01:00',
        CEndTime: '02:00',
        CDeltaTime: '01:00',
        FEndDrillDepth: 1000,
        FEndToolDepth: 1000,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '02:00',
        CEndTime: '04:20',
        CDeltaTime: '02:20',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'бурение',
        VcGroupOper: 'бурение',
        VcComment: 'Бурение в инт. 1100-1240м',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      },
      {
        VcDate: '11.11.19',
        CBegTime: '04:20',
        CEndTime: '06:30',
        CDeltaTime: '02:10',
        FEndDrillDepth: 1240,
        FEndToolDepth: 1240,
        CZatrName: 'Промывка',
        VcGroupOper: 'бурение',
        VcComment: 'Промывка на забое',
        IRangeNum: 5,
        IStageNum: 2,
        ICasingNum: 2,
        VcSolutions: ''
      }
    ]
  })
}
</script>
