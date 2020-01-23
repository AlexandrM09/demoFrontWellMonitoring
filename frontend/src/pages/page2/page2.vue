<template>
<div class="title" >Протокол
      <v-data-table
      class=" blue-grey lighten-5 table table-bordered table-hover table-sm "  border="2" height="10%" with="100%"
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
        align: 'centr',
        sortable: false,
        value: 'VcDate',
        class: colorHeader,
        width: '20px'
      },
      { text: 'Начало', value: 'CBegTime', class: colorHeader, width: 40 },
      { text: 'Конец', value: 'CEndTime', class: colorHeader },
      { text: 'Всего', value: 'CDeltaTime', class: colorHeader },
      { text: 'Забой', value: 'FEndDrillDepth', class: colorHeader },
      { text: 'Пол. долота', value: 'FEndToolDepth', class: colorHeader },
      { text: 'Код', value: 'CZatrName', class: colorHeader },
      { text: 'Операция', value: 'VcGroupOper', class: colorHeader },
      { text: 'Описание работ', value: 'VcComment', class: colorHeader, width: 200 },
      { text: 'Рейс', value: 'IRangeNum', class: colorHeader },
      { text: 'Этап', value: 'IStageNum', class: colorHeader },
      { text: 'Ок', value: 'ICasingNum', class: colorHeader },
      { text: 'Парам. р-ра', value: 'VcSolutions', class: colorHeader }
    ],
    protocol: [
      {
        VcDate: '11.11.19',
        CBegTime: '00:00',
        CEndTime: '01:00',
        CDeltaTime: '01:00',
        FEndDrillDepth: 100,
        FEndToolDepth: 100,
        CZatrName: 'бурение',
        VcGroupOper: 'бурение',
        VcComment: 'Бурение в инт. 50-100м',
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
        FEndDrillDepth: 100,
        FEndToolDepth: 100,
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
