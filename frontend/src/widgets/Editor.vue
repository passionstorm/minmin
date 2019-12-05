<style>
  .editor .ProseMirror {
    box-shadow: inset 0 1px 2px rgba(10, 10, 10, 0.1);
    max-width: 100%;
    width: 100%;
    background-color: white;
    border-color: #dbdbdb;
    border-radius: 4px;
    color: #363636;
    padding: 10px;
  }

  .editor .menubar {
    display: flex;
    height: 150px;
    -webkit-transition: visibility .2s .4s, opacity .2s .4s;
    transition: visibility .2s .4s, opacity .2s .4s;
  }

  .editor .ProseMirror p {
    margin: 0;
  }

  .editor .btn_menubar {
    font-weight: 700;
    display: inline-flex;
    background: transparent;
    border: 0;
    color: rgba(0, 0, 0, .54);
    margin-left: 8px;
    border-radius: 3px;
    cursor: pointer;
    text-decoration: none;
    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -moz-user-select: none; /* Old versions of Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none;
  }

  .editor .w {
    border-left: 1px solid #d3d9df;
    height: 20px;
    line-height: normal;
    list-style: none;
    margin-left: 10px;
    outline: none;
    overflow: hidden;
    padding: 0;
    text-decoration: none;
    vertical-align: middle;
    width: 0;
  }

  .editor .text_toolbar {
    align-items: center;
    border-radius: 2px;
    -webkit-box-shadow: 0 4px 5px 0 rgba(0, 0, 0, 0.14), 0 1px 10px 0 rgba(0, 0, 0, 0.12), 0 2px 4px -1px rgba(0, 0, 0, 0.2);
    box-shadow: 0 4px 5px 0 rgba(0, 0, 0, 0.14), 0 1px 10px 0 rgba(0, 0, 0, 0.12), 0 2px 4px -1px rgba(0, 0, 0, 0.2);
    display: inline-flex;
    margin: 4px 0;
    padding: 4px 2px;
    flex-wrap: wrap;
    -webkit-flex-wrap: wrap;
    white-space: nowrap;
  }

  .editor .toolbar {
    display: flex;
    flex-direction: column;
  }

  .editor .btn_menubar.is-active {
    background-color: #dddddd;
  }

  .editor .ext_toolbar {
    display: flex;
    justify-content: space-between;
    flex-wrap: wrap;
    -webkit-flex-wrap: wrap;
    padding: 0 10px;
  }

  .editor .btn_ext_toolbar.active {
    background: #dddddd;
  }

  .editor .btn_ext_toolbar {
    display: inline-flex;
    border-radius: 15px;
    padding: 4px 8px;
    cursor: pointer;
    justify-content: center;
    text-decoration: none;
    position: relative;
    align-items: center;
    background-color: #f5f5f5;
    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -moz-user-select: none; /* Old versions of Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none;
    margin: 4px;
  }

  /*.editor .btn_ext_toolbar:not(:first-child) {*/
  /*  margin-left: 8px;*/
  /*}*/

  .editor .btn_ext_toolbar .ic {
    width: 20px;
    height: 20px;
    display: flex;
  }

  .editor .btn_ext_toolbar .tl {
    padding: 0 10px 0 8px;
    display: flex;
  }

  ul[data-type="todo_list"] {
    padding-left: 0;
  }

  li[data-type="todo_item"] {
    display: flex;
    flex-direction: row;
  }

  .todo-checkbox {
    width: 1.25em;
    height: 1.25em;
    flex-shrink: 0;
    border-radius: 4px;
    border: 2px solid #7a7a7a;
    transition: background 150ms ease-out;
    background: transparent;
    user-select: none;
    -webkit-user-select: none;
    cursor: pointer;
  }

  li[data-done="true"] .todo-content .p {
    text-decoration: line-through;
  }

  li[data-done="true"] .todo-checkbox {
    background: #80bdff url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 1 1'%3E%3Cpath style='fill:white' d='M 0.04038059,0.6267767 0.14644661,0.52071068 0.42928932,0.80355339 0.3232233,0.90961941 z M 0.21715729,0.80355339 0.85355339,0.16715729 0.95961941,0.2732233 0.3232233,0.90961941 z'%3E%3C/path%3E%3C/svg%3E") no-repeat center center;
    border-color: #80bdff;
  }

  .editor_content table {
    border-collapse: collapse;
    table-layout: fixed;
    width: 100%;
    margin: 0;
    overflow: hidden;
  }

  .editor_content table th {
    font-weight: 700;
    text-align: left;
  }

  .editor_content table th,
  .editor_content table td {
    min-width: 1em;
    border: 2px solid #ddd !important;
    padding: 3px 5px;
    vertical-align: top;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    position: relative;
  }

  .content table tbody tr:last-child td,
  .content table tbody tr:last-child th {
    border-bottom-width: 1px;
  }
  .editor_content table .selectedCell:after {
    z-index: 2;
    position: absolute;
    content: "";
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    background: rgba(200, 200, 255, .4);
    pointer-events: none
  }

  .editor_content table .column-resize-handle {
    position: absolute;
    right: -2px;
    top: 0;
    bottom: 0;
    width: 4px;
    z-index: 20;
    background-color: #adf;
    pointer-events: none
  }

  .editor_content table td > *, .editor_content table th > * {
    margin-bottom: 0
  }

  .editor_content img {
    max-width: 100%;
    border-radius: 3px
  }

  .m-footer {
    width: 100%;
    position: fixed;
    bottom: 20px;
    max-width: 960px;
    background: #fff;
  }

  .text_toolbar, .menubar, .ext_toolbar {
    width: 100%;
  }

  .modal-content {
    background: #fff;
    max-width: 960px !important;
    height: calc(100vh - 218px);
    max-height: 100vh;
  }

  .modal {
    justify-content: inherit;
    padding-top: 3em;
  }

  .w-r {
    border: 1px solid #ccc;
    box-shadow: 0 -5px 5px -5px #333;
  }
</style>

<template>
  <div class="editor" @input="onInput"
       @blur="onBlur"
       @focus="onFocus">
    <div @click="modalShow = true" class="input">{{editor ? editor.getHTML() : '' | truncate('50','...')}}</div>
    <modal :active.sync="modalShow" :width="640" scroll="keep">
      <div style="height: 100%" @click="editor.focus()">
        <editor-content :class="{has_chil_bar: hasChildBar}" class="editor_content" :editor="editor"/>
      </div>
      <div slot="footer" class="m-footer">
        <editor-menu-bar class="menubar w-r" :editor="editor" v-slot="{ commands, isActive }">
          <div class="toolbar">
            <div class="text_toolbar" v-show="textBar">
              <a @click="commands.bold" class="btn_menubar" :class="{ 'is-active': isActive.bold() }">
                <icon name="format_bold"/>
              </a>
              <a class="btn_menubar" :class="{ 'is-active': isActive.italic() }" @click="commands.italic">
                <icon name="format_italic"/>
              </a>
              <a class="btn_menubar" :class="{ 'is-active': isActive.strike() }" @click="commands.strike">
                <icon name="format_clear"/>
              </a>
              <a class="btn_menubar" :class="{ 'is-active': isActive.underline() }" @click="commands.underline">
                <icon name="underline"/>
              </a>
              <div class="w"></div>
              <a class="btn_menubar" :class="{ 'is-active': isActive.bullet_list() }" @click="commands.bullet_list">
                <icon name="format_list_bulleted"/>
              </a>
              <a class="btn_menubar" :class="{ 'is-active': isActive.ordered_list() }" @click="commands.ordered_list">
                <icon name="format_list_numbered"/>
              </a>
              <a class="btn_menubar" :class="{ 'is-active': isActive.code_block() }" @click="commands.code_block">
                <icon name="code-24px"/>
              </a>
              <div class="w"></div>
              <a class="btn_menubar" :class="{ 'is-active': isActive.todo_list() }" @click="commands.todo_list">
                <icon name="task"/>
              </a>
              <a class="btn_menubar" :class="{ 'is-active': isActive.table() }"
                 @click="commands.createTable({rowsCount: 3, colsCount: 3, withHeaderRow: false })">
                <icon name="table"/>
              </a>
              <div style="margin: auto"></div>
              <a class="btn_menubar" @click="commands.undo">
                <icon name="undo"/>
              </a>
              <a class="btn_menubar" @click="commands.redo">
                <icon name="redo"/>
              </a>
            </div>
            <div class="text_toolbar" v-show="onHasChildBar(isActive.table())">
              <a @click="commands.deleteTable" class="btn_menubar">
                <icon name="table_del"/>
              </a>
              <a @click="commands.addColumnBefore" class="btn_menubar">
                <icon name="add_col_l"/>
              </a>
              <a @click="commands.addColumnAfter" class="btn_menubar">
                <icon name="add_col_r"/>
              </a>
              <a @click="commands.deleteColumn" class="btn_menubar">
                <icon name="del_col"/>
              </a>
              <a @click="commands.addRowBefore" class="btn_menubar">
                <icon name="add_row_t"/>
              </a>
              <a @click="commands.addRowAfter" class="btn_menubar">
                <icon name="add_row_b"/>
              </a>
              <a @click="commands.deleteRow" class="btn_menubar">
                <icon name="del_row"/>
              </a>
            </div>
            <div class="ext_toolbar">
              <div class="btn_ext_toolbar" :class="{active: textBar}" @click="textBar = !textBar">
                <icon class="ic" name="text_format" @click="commands.image"/>
                <div class="tl">Công cụ</div>
              </div>
              <div class="btn_ext_toolbar">
                <icon class="ic" name="insert_photo" @click="commands.image"/>
                <div class="tl">Ảnh</div>
              </div>
              <div class="btn_ext_toolbar">
                <icon name="link" @click="commands.link"/>
                <div class="tl">Chèn link</div>
              </div>
              <div class="btn_ext_toolbar">
                <icon name="grin" @click="commands.link"/>
                <div class="tl">Cảm xúc</div>
              </div>
            </div>
          </div>
        </editor-menu-bar>
      </div>
    </modal>

  </div>
</template>

<script>
  import {
    Blockquote,
    Bold,
    BulletList,
    Code,
    CodeBlock,
    HardBreak,
    Heading,
    History,
    HorizontalRule,
    Image,
    Italic,
    Link,
    ListItem,
    OrderedList,
    Strike,
    Table,
    TableCell,
    TableHeader,
    TableRow,
    TodoItem,
    TodoList,
    Underline,
  } from 'tiptap-extensions';
  import {Editor, EditorContent, EditorMenuBar} from 'tiptap';
  import {Icon, Modal} from './';

  const EXTENSIONS = [
    new Blockquote(),
    new BulletList(),
    new HardBreak(),
    new Heading({levels: [1, 2, 3]}),
    new HorizontalRule(),
    new ListItem(),
    new TodoList(),
    new Table({
      resizable: true,
    }),
    new TableHeader(),
    new TableCell(),
    new TableRow(),
    new TodoItem({
      nested: true,
    }),
    new TodoList(),
    new OrderedList(),
    new Link(),
    new Bold(),
    new Code(),
    new Italic(),
    new Strike(),
    new CodeBlock(),
    new Underline(),
    new History(),
    new Image(),
  ];

  export default {
    name: 'editor',
    components: {
      EditorContent,
      EditorMenuBar,
      Modal,
      Icon,
    },
    props: {
      value: [Number, String],
    },
    data() {
      return {
        hasChildBar: false,
        modalShow: false,
        editor: null,
        textBar: true,
      };
    },
    mounted() {
      this.editor = new Editor({
        extensions: EXTENSIONS,
        onFocus: this.onFocus,
        onUpdate: this.onUpdate,
        content: this.value,
      });
    },
    computed: {},
    methods: {
      onHasChildBar(val) {
        this.hasChildBar = val;
        return this.hasChildBar;
      },
      onInput(event) {
        this.$nextTick(() => {
          if (event.target) {
            this.computedValue = event.target.value;
          }
        });
      },
      onBlur(val) {
        this.$emit('blur', val);
      },
      onFocus(val) {
        this.modalShow = true;
        this.$emit('focus', val);
      },
      onUpdate(val) {
        this.$emit('input', val.getHTML());
      },
    },
    beforeDestroy() {
      this.editor.destroy();
    },
  };
</script>