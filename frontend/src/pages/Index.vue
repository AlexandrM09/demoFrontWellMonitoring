<template>
 <div id="app">
  <v-app id="inspire">
    <v-navigation-drawer
    src="static/29157-elektrik-akva-biryuza-svet-sinij-2560x1440.jpg"
    clipped
    app
    class="blue lighten-1 "
    v-model="drawer"
    dark
    width=250
    >
      <v-list dense>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="{path: '/page' + (i+1)}"
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title> {{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar app fixed clipped-left height="48px" color="light-blue lighten-5">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer "></v-app-bar-nav-icon>
      <v-toolbar-title>Скважина: </v-toolbar-title>
          <v-menu
          left
          bottom
        >
          <template v-slot:activator="{ on }" centr>
            <v-btn icon v-on="on" centr>
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
          </template>
            <v-list >
            <v-list-item
              v-for="n in 5"
              :key="n"
              @click="() => {}"
            >
              <v-list-item-title>Скважина {{ n }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
    </v-app-bar >
    <v-content >
      <v-container fluid  >
        <v-fade-transition mode="out-in">
          <v-img
        height="100%"
        src="static/4c78472f-f7ca-4fc6-80c9-fd925cd8a72f.png"
      >
          <router-view></router-view>
          </v-img>
        </v-fade-transition>
      </v-container>
    </v-content>
    </v-app>
  </div>
</template>
<style App.vue>
.v-data-table th {
  font-size: 14px;
  color: #727dbd;
  vertical-align: bottom;
}

.v-data-table td {
  font-size: 6x;
}
</style>>

<script>
// static/arka.jpg
// src="static/linii_svet_yarkiy_figura_47031_1366x768.jpg"
import 'material-design-icons-iconfont/dist/material-design-icons.css'
export default {
  data: () => ({
    drawer: null,
    items: [
      {icon: 'timeline', text: 'Графики'},
      {icon: 'view_list', text: 'Протокол'},
      {icon: 'trending_up', text: 'График строительства'},
      {icon: 'notes', text: 'Сведения о скважине'},
      {icon: 'pie_chart', text: 'Баланс времени'},
      {icon: 'insert_chart_outlined', text: 'Профиль скважины'}
    ]
  }),
  props: {
    source: String
  }
}
</script>
