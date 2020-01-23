<template>
<div class="title" >Протокол
      <v-data-table
      class=" table table-bordered table-hover table-sm "  border="2" height="10%" with="100%"
      :headers="headers"
      :items="protocol"
      :items-per-page="30"
      :rows-per-page-items='[30,10,{text: "Все", value: -1}]'
      >
      </v-data-table>
</div>
</template>
<script>

export default {
  data: () => ({
    headers: [
      {
        text: 'Дата',
        align: 'left',
        sortable: false,
        value: 'VcDate'
      },
      { text: 'Начало', value: 'CBegTime' },
      { text: 'Конец', value: 'CEndTime' },
      { text: 'Всего', value: 'CDeltaTime' },
      { text: 'Забой', value: 'FEndDrillDepth' },
      { text: 'Пол. долота', value: 'FEndToolDepth' },
      { text: 'Код', value: 'CZatrName' },
      { text: 'Операция', value: 'VcGroupOper' },
      { text: 'Описание работ', value: 'VcComment' },
      { text: 'Рейс', value: 'IRangeNum' },
      { text: 'Этап', value: 'IStageNum' },
      { text: 'Ок', value: 'ICasingNum' },
      { text: 'Парам. р-ра', value: 'VcSolutions' }
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
