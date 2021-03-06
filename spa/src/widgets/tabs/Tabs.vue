<template>
  <div class="b-tabs" :class="mainClasses">
    <nav class="tabs" :class="navClasses">
      <ul>
        <li
            v-for="(tabItem, index) in tabItems"
            :key="index"
            v-show="tabItem.visible"
            :class="{ 'is-active': activeTab === index, 'is-disabled': tabItem.disabled }">
          <a @click="tabClick(index)">
            <template v-if="tabItem.$slots.header">
              <b-slot-component
                  :component="tabItem"
                  name="header"
                  tag="span"/>
            </template>
            <template v-else>
              <icon
                  v-if="tabItem.icon"
                  :name="tabItem.icon"
                  :pack="tabItem.iconPack"
                  :size="size"/>
              <span>{{ tabItem.label }}</span>
            </template>
          </a>
        </li>
      </ul>
    </nav>
    <section class="tab-content" :class="{'is-transitioning': isTransitioning}">
      <slot/>
    </section>
  </div>
</template>

<script>
  import {Icon} from '../';
  import SlotComponent from '../mixins/slot.mixin';

  export default {
    name: 'Tabs',
    components: {
      Icon,
      SlotComponent,
    },
    props: {
      value: Number,
      expanded: Boolean,
      type: String,
      size: String,
      position: String,
      animated: {
        type: Boolean,
        default: true,
      },
      destroyOnHide: {
        type: Boolean,
        default: false,
      },
      vertical: Boolean,
    },
    data() {
      return {
        activeTab: this.value || 0,
        tabItems: [],
        contentHeight: 0,
        isTransitioning: false,
        _isTabs: true, // Used internally by TabItem
      };
    },
    computed: {
      mainClasses() {
        return {
          'is-fullwidth': this.expanded,
          'is-vertical': this.vertical,
          [this.position]: this.position && this.vertical,
        };
      },
      navClasses() {
        return [
          this.type,
          this.size,
          {
            [this.position]: this.position && !this.vertical,
            'is-fullwidth': this.expanded,
            'is-toggle-rounded is-toggle': this.type === 'is-toggle-rounded',
          },
        ];
      },
    },
    watch: {
      /**
       * When v-model is changed set the new active tab.
       */
      value(value) {
        this.changeTab(value);
      },
      /**
       * When tab-items are updated, set active one.
       */
      tabItems() {
        if (this.activeTab < this.tabItems.length) {
          this.tabItems[this.activeTab].isActive = true;
        }
      },
    },
    methods: {
      /**
       * Change the active tab and emit change event.
       */
      changeTab(newIndex) {
        if (this.activeTab === newIndex || this.tabItems[newIndex] === undefined) return;
        if (this.activeTab < this.tabItems.length) {
          this.tabItems[this.activeTab].deactivate(this.activeTab, newIndex);
        }
        this.tabItems[newIndex].activate(this.activeTab, newIndex);
        this.activeTab = newIndex;
        this.$emit('change', newIndex);
      },
      /**
       * Tab click listener, emit input event and change active tab.
       */
      tabClick(value) {
        if (this.activeTab === value) return;
        this.$emit('input', value);
        this.changeTab(value);
      },
    },
    mounted() {
      if (this.activeTab < this.tabItems.length) {
        this.tabItems[this.activeTab].isActive = true;
      }
    },
  };
</script>