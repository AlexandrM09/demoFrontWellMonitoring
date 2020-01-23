<template>
<div class="title">Сведения о скважине
<v-layout row wrap>
  <v-flex d-flex  space-around xs12 sm8 md5>
  <v-card
      class="left mx-auto"
      elevation=5
    >
    <v-simple-table fixed-header >
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Параметр</th>
            <th class="text-left">Значение</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.name">
            <td>{{ item.name }}</td>
            <td>{{ item.data }}</td>
          </tr>
        </tbody>
      </template>
     </v-simple-table>
    </v-card>
    </v-flex>
    <v-flex d-flex xs12 sm6 md4>
      <v-card
      class=""
      elevation=5
    >
    <v-simple-table fixed-header >
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Параметр</th>
            <th class="text-left">Значение</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.name">
            <td>{{ item.name }}</td>
            <td>{{ item.data }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    </v-card>
    </v-flex>
 </v-layout>
</div>
</template>
<script>

let colorHeader = 'blue lighten-4'
export default {
  data: () => ({
    items: [ {name: 'Имя скважины', data: 'Уренгойская'},
      {name: 'Номер скважины', data: '210'},
      {name: '№ Куста', data: '2'},
      {name: '№ скв', data: ''},
      {name: 'Партия', data: '24'},
      {name: 'Начальник партии', data: 'Иванов А.А.'},
      {name: 'Тип скважины', data: ''},
      {name: 'Назначание скважины', data: ''},
      {name: 'Площадь', data: ''},
      {name: 'Вращение инструмента', data: ''},
      {name: 'Месторождение', data: ''}
    ],
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
