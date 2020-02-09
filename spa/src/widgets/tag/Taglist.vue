<template>
  <div class="tags" :class="{ 'has-addons': attached }">
    <slot/>
  </div>
</template>

<script>
  export default {
    name: 'Taglist',
    props: {
      attached: Boolean
    }
  }
</script>
