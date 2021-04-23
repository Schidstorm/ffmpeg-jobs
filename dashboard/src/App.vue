<template>
  <v-app>
    <v-app-bar
      app
      dark
    >
      <v-img
        alt="Vuetify Logo"
        class="shrink mr-2"
        contain
        :src="require('./assets/logo.svg')"
        transition="scale-transition"
        width="30"
      />
      <v-app-bar-title class="title">
        Ffmpeg Jobs Dashboard
      </v-app-bar-title>
    </v-app-bar>

    <v-main>
      <JobList :apiServer="this.$data.config.apiServerUrl"/>
    </v-main>
  </v-app>
</template>

<script>
import JobList from './components/JobList';
import * as config from './config'

export default {
  name: 'App',

  components: {
    JobList,
  },

  data: () => {
    return {
      config: config.current
    }
  },
};
</script>

<style>
.title .v-app-bar-title__content {
  text-overflow: unset;
}
</style>