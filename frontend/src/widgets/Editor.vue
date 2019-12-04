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
    /*background-color: #f5f5f5 !important;*/
    /*border-color: #f5f5f5 !important;*/
    display: flex;
    height: auto !important;
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
    border-left: 1px solid rgba(100, 121, 143, 0.122);
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
    margin: 5px 0;
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
    margin-left: 8px;
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
</style>

<template>
  <div class="editor">
    <editor-content class="editor_content" :editor="editor"/>
    <editor-menu-bar class="menubar" :editor="editor" v-slot="{ commands, isActive }">
      <div class="toolbar">
        <div class="text_toolbar" :style="{visibility: textBar ? 'visible' : 'hidden'}">
          <a @click="commands.bold" class="btn_menubar" :class="{ 'is-active': isActive.bold() }">
            <icon name="format_bold"/>
          </a>
          <a class="btn_menubar" :class="{ 'is-active': isActive.italic() }" @click="commands.italic">
            <icon name="format_italic"/>
          </a>
          <a class="btn_menubar" :class="{ 'is-active': isActive.strike() }" @click="commands.strike">
            <icon name="format_italic"/>
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
            <icon name="format_list_numbered"/>
          </a>
          <div style="margin: auto"></div>
          <a class="btn_menubar" @click="commands.undo">
            <icon name="undo"/>
          </a>
          <a class="btn_menubar" @click="commands.redo">
            <icon name="redo"/>
          </a>
        </div>
        <div class="ext_toolbar">
          <div class="btn_ext_toolbar" :class="{active: textBar}" @click="textBar = !textBar">
            <icon class="ic" name="text_format" @click="commands.image"/>
            <div class="tl"></div>
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
    Underline,
  } from 'tiptap-extensions';
  import {Editor, EditorContent, EditorMenuBar} from 'tiptap';
  import {Icon} from './';

  const EXTENSIONS = [
    new Blockquote(),
    new BulletList(),
    new HardBreak(),
    new Heading({levels: [1, 2, 3]}),
    new HorizontalRule(),
    new ListItem(),
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
      Icon,
    },
    data() {
      return {
        editor: new Editor({
          extensions: EXTENSIONS,
          content: `
          <h2>
            Hi there,
          </h2>
          <p>
            this is a very <em>basic</em> example of tiptap.
          </p>
          <pre><code>body { display: none; }</code></pre>
          <ul>
            <li>
              A regular list
            </li>
            <li>
              With regular items
            </li>
          </ul>
          <blockquote>
            It's amazing 👏
            <br />
            – mom
          </blockquote>
        `,
        }),
        textBar: true,
      };
    },
    beforeDestroy() {
      this.editor.destroy();
    },
  };
</script>