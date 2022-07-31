<template>
  <el-container>
    <el-header>
      <el-row>
        <div class="logo-field" v-on:click="doLogoLink">
          <img class="logo-img" src="@/assets/logo.png"/>
          <p class="proj-name">Takoyaki</p>
        </div>
      </el-row>
    </el-header>
    <el-container>
      <el-aside>
        <el-row v-for="(item, idx) in menus" :key="idx">
          <p class="menu-item" :class="{'menu-item-active': idx === activeMenu}" @click="switchMenuFocus(item, idx)">
            <i :class="item.icon"></i> {{item.name}}
          </p>
        </el-row>
      </el-aside>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      // 菜单导航
      activeMenu: 0,
      menus: [
        {
          path: "/database",
          icon: "el-icon-menu",
          name: "数据源"
        },
        {
          path: "/page",
          icon: "el-icon-tickets",
          name: "页面配置"
        }
      ]
    }
  },
  computed: {
    activeMenuTitle() {
      return this.menus[this.activeMenu].name
    }
  },
  methods: {
    // 菜单切换
    switchMenuFocus(item, idx) {
      this.activeMenu = idx
      this.$router.push(item.path)
    },
    backToHomePage() {
      this.$router.push('/')
    },
    doLogoLink() {
      window.location.href = '/'
    }
  }
}
</script>

<style lang="scss">
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
}
.el-container {
  height: 100%;
}
.el-header {
  padding: 0 !important;
  background-color: #262f3e;
  .logo-field {
    width: 169px;
    border-right: 1px solid #000000;
    display: flex;
    align-items: center;
    padding: 0 15px;
    .proj-name{
      color: #ffffff;
      font-size: 20px;
      font-weight: bold;
      margin: 0;
      line-height: 60px;
      text-align: center;
      text-indent: 4px;
    }
    .logo-img {
      width: 44px;
      height: 44px;
    }
  }
  .logo-field:hover{
    cursor: pointer;
  }
}
.el-aside {
  background-color: #1f222c;
  width: 200px !important;
  .menu-item {
    color: #c1c6c8;
    text-align: left;
    padding: 15px 20px;
    margin: 0;
  }
  .menu-item:hover{
    background-color: #2d3546;
    color: #ffffff;
    cursor: pointer;
  }
  .menu-item-active {
    background-color: #2d3546;
    color: #ffffff;
    cursor: pointer;
  }
}
.el-main {
  padding: 0 !important;
  background-color: #f3f4f7;
  .content-header {
    background-color: #ffffff;
    .page-title {
      padding-left: 20px;
      text-align: left;
      font-weight: bold;
    }
  }
  .content-body {
    text-align: left;
    padding: 20px;
  }
}
.el-input__count {
  line-height: 20px;
}
</style>