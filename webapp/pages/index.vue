<template>
  <h3>{{ user?.name }}</h3>
</template>

<script>
export default {
  name: 'IndexPage',
  middleware: ['auth'],
  computed: {
    user () {
      return this.$auth?.user
    }
  }
}
</script>
