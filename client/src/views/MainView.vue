<script setup>
import TheWelcome from '../components/TheWelcome.vue'
</script>

<template>
  <div class="">
    <div class="bg-sky-700 m-4 p-4 flex justify-center flex-col rounded-lg">
      <div class="flex flex-col justify-around mx-4 mb-4">
        <h2 class="text-3xl text-white">Hello</h2>
        <h3 class="text-xl text-white">Welcome to template page!</h3>
      </div>
      <button
        type="button"
        class="relative inline-flex items-center justify-center p-0.5 mb-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500 group-hover:from-cyan-500 group-hover:to-blue-500 hover:text-white dark:text-white"
        @click="submit()"
      >
        <span
          class="w-full relative px-5 py-2.5 transition-all ease-in duration-75 bg-gray-900 rounded-md group-hover:bg-opacity-0"
        >
          Test Server Respose
        </span>
      </button>
      <TheWelcome class="rounded-lg" />
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      username: 'Gunta'
    }
  },

  methods: {
    submit() {
      axios
        .post('http://localhost:8080/', { Username: 'Gunta' })
        .then((res) => {
          console.log(res.data)
        })
        .catch((error) => {
          console.error('Error at login data fetching: ' + error)
        })
    }
  }
}
</script>
