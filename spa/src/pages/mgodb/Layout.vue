<template>
  <div>
    <pre>
      {{dt}}
    </pre>
    <pre class="json-viewer">
      <a class="list-link" href="javascript:void(0)">{<span class="items-ph hide">6 items</span>
      </a>
      <ul data-level="0" class="type-object">
        <li v-for="i in dt">
          <a href="" class="list-link">{{i}}</a>
          <ul data-level="2" class="type-object hide">
            <li>d: <span class="type-number">4</span></li>
          </ul>
        </li>
      </ul>
      }
    </pre>


  </div>
</template>

<script>
  export default {
    name: 'index.vue',
    data() {
      return {
        dt: {},
        collects: {},
      };
    },
    methods: {
      onGetCollect(name) {
        let t = this;
        axios.get('/db/' + name).then(rs => {
          t.collects[name] = rs.data;
        });
      },
    },
    computed: {},
    created() {
      let t = this;
      axios.get('/db').then(rs => {
        t.dt = rs.data;
      });
    },
  };
</script>

<style lang="scss" scoped>
  .json-viewer {
    color: #000;
    display: block;
    font-family: monospace;
    white-space: pre;
    margin: 1em 0px;
    padding-left: 20px;
    a.list-link:before {
      color: #aaa;
      content: "\25BC";
      position: absolute;
      display: inline-block;
      width: 1em;
      left: -1em;
    }
    ul {
      list-style-type: none;
      margin: 0;
      margin: 0 0 0 1px;
      border-left: 1px dotted #ccc;
      padding-left: 2em;

      li {

      }
    }

    .hide {
      display: none;
    }
  }


  .json-viewer ul li .type-string,
  .json-viewer ul li .type-date {
    color: #0B7500;
  }

  .json-viewer ul li .type-boolean {
    color: #1A01CC;
    font-weight: bold;
  }

  .json-viewer ul li .type-number {
    color: #1A01CC;
  }

  .json-viewer ul li .type-null {
    color: red;
  }

  .json-viewer a.list-link {
    color: #000;
    text-decoration: none;
    position: relative;
  }

  .json-viewer a.list-link:before {
    color: #aaa;
    content: "\25BC";
    position: absolute;
    display: inline-block;
    width: 1em;
    left: -1em;
  }

  .json-viewer a.list-link.collapsed:before {
    content: "\25B6";
  }

  .json-viewer a.list-link.empty:before {
    content: "";
  }

  .json-viewer .items-ph {
    color: #aaa;
    padding: 0 1em;
  }

  .json-viewer .items-ph:hover {
    text-decoration: underline;
  }

</style>