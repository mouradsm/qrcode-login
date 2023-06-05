<template>
  <div class="flex h-screen">
    <div class="max-w-xs m-auto">
      <form
        class="px-8 pt-6 pb-8 mb-4 bg-white rounded shadow-md"
        @submit.prevent="userLogin"
      >
        <div class="mb-4">
          <label
            class="block mb-2 text-sm font-bold text-gray-700"
            for="username"
          >
            Username
          </label>
          <input
            id="username"
            v-model="login.email"
            class="w-full px-3 py-2 leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
            type="text"
            placeholder="Username"
          >
        </div>
        <div class="mb-6">
          <label
            class="block mb-2 text-sm font-bold text-gray-700"
            for="password"
          >
            Password
          </label>
          <input
            id="password"
            v-model="login.password"
            class="w-full px-3 py-2 mb-3 leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
            type="password"
            placeholder="******************"
          >
          <!-- <p class="text-xs italic text-red-500">
          Please choose a password.
        </p> -->
        </div>
        <div class="flex flex-col items-center justify-between">
          <button
            class="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Sign In
          </button>
          <a
            class="inline-block mt-5 text-xs font-bold text-blue-500 align-baseline hover:text-blue-800"
            href="#"
          >
            Forgot Password?
          </a>
        </div>
      </form>
      <!-- <p class="text-xs text-center text-gray-500">
        &copy;2020 Acme Corp. All rights reserved.
      </p> -->
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      login: {
        email: '',
        password: ''
      },
      error: null
    }
  },
  methods: {
    async userLogin () {
      try {
        await this.$auth.loginWith('local', {
          data: this.login
        })

        console.log(this.$auth.user)

        this.$router.push('/')
      } catch (err) {
        this.error = err.response.data.message
      }
    }
  }
}
</script>
