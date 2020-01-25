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
          <tr v-for="item in items" :key="item.name" class="cyan lighten-5">
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
        <thead >
          <tr class="cyan darken-3">
            <th class="text-left">Параметр</th>
            <th class="text-left">Значение</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items2" :key="item.name" class="cyan lighten-5">
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

// let colorHeader = 'blue lighten-4'
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
    items2: [ {name: 'Операторы', data: 'Иванов А.А., Петров В.В.'},
      {name: 'Начало строительство', data: '19.12.19'},
      {name: 'Окончание строит. план', data: '23.01.20'},
      {name: 'Начало ГТИ', data: ''},
      {name: 'Окончание ГТИ факт', data: ''},
      {name: 'Буровая установка', data: 'Иванов А.А.'},
      {name: 'Тип насосов', data: ''},
      {name: 'Тип ротора(СВП)', data: ''},
      {name: 'Тип превентора', data: ''},
      {name: 'Тип дегазатора', data: ''}
    ]
  })
}
</script>
