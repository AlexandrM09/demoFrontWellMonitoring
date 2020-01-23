<template>
 <div id="app">
  <v-app id="inspire">
    <v-navigation-drawer app
    class="blue darken-3 "
    v-model="drawer"
    dark
    clipped
     permanent>
      <v-list dense>
        <v-list-item
          v-for="i in 4"
          :key="i"
          :to="{path: '/page' + i}"
        >
          <v-list-item-action>
            <v-icon>{{ icons[i-1] }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title> {{ pagenames[i-1] }}</v-list-item-title>
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
              <v-list-item-title>Option {{ n }}</v-list-item-title>
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
  font-size: 12x;
}
</style>>

<script>
import 'material-design-icons-iconfont/dist/material-design-icons.css'
export default {
  data: () => ({
    drawer: true,
    icons: ['home', 'view_list', 'trending_up', 'notes'],
    pagenames: ['Графики', 'Протокол', 'График строительства', 'Сведения о скважине']
  }),
  props: {
    source: String
  }
}
</script>
