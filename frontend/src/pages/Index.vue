<template>
 <div id="app">
  <v-app id="inspire">
    <v-navigation-drawer
    permanent
    clipped
    app
    class="blue lighten-1 "
    v-model="drawer"
    dark
    width=200
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
    <v-app-bar app clipped-left height="48px">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer "></v-app-bar-nav-icon>
      <v-toolbar-title>Скважина: </v-toolbar-title>
          <v-menu
          left
          bottom
        >
          <template v-slot:activator="{ on }" right>
            <v-btn icon v-on="on" right>
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
    <v-content>
      <v-container fluid>
        <v-fade-transition mode="out-in">
          <router-view></router-view>
        </v-fade-transition>
      </v-container>
    </v-content>
    </v-app>
  </div>
</template>
<style App.vue>
.v-data-table th {
  font-size: 14px;
  color: #727dbd;;
}

.v-data-table td {
  font-size: 6x;
}
</style>>

<script>
import 'material-design-icons-iconfont/dist/material-design-icons.css'
export default {
  data: () => ({
    drawer: null,
    items: [
      {icon: 'home', text: 'Графики'},
      {icon: 'view_list', text: 'Протокол'},
      {icon: 'trending_up', text: 'График строительства'},
      {icon: 'notes', text: 'Сведения о скважине'}
    ],
    icons: ['home', 'view_list', 'trending_up', 'notes'],
    pagenames: ['Графики', 'Протокол', 'График строительства', 'Сведения о скважине']
  }),
  props: {
    source: String
  }
}
</script>
