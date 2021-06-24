<template>
  <div class="big">
    <el-row>
      <div class="card">
        <el-col :xs="24" :lg="16" :md="16">
          <div class="car-left">
            <el-row>
              <div>
                <el-col :xs="4" :md="3" :lg="3">
                  <span class="card-img">
                    <img :src="userInfo.headerImg" alt="">
                  </span>
                </el-col>
                <el-col :xs="20" :lg="12" :md="12">
                  <div class="text">
                    <h4>早安，今天也是充满爆肝的一天</h4>
                    <p class="tips-text">
                      <i class="el-icon-sunny el-icon" />
                      <span>今日多云，26℃ - 38℃，天气多变，出门记得带伞。</span>
                    </p>
                  </div>
                </el-col>
              </div>
            </el-row>
          </div>
        </el-col>
        <el-col :xs="24" :lg="8" :md="8">
          <div class="car-right">
            <el-row>
              <el-col :span="8">
                <div class="car-item">
                  <span class="flow"><i class="el-icon-s-grid" /></span>
                  <span>今日流量 </span>
                  <b>14</b>
                </div>
              </el-col>
              <el-col :span="8">
                <div class="car-item">
                  <span class="user-number">
                    <i class="el-icon-s-custom" />
                  </span>
                  <span>总用户 </span>
                  <b>2</b>
                </div>
              </el-col>
              <el-col :span="8">
                <div class="car-item">
                  <span class="feedback">
                    <i class="el-icon-star-on" />
                  </span>
                  <span>好评率 </span>
                  <b>100%</b>
                </div>
              </el-col>
            </el-row>
          </div>
        </el-col>
      </div>
    </el-row>

    <el-row :gutter="20">
      <el-col :span="12" class="shadow">
        <el-row :gutter="20">
          <el-col
            v-for="(card, key) in toolCards"
            :key="key"
            :span="8"
            :sm="8"
            :push="1"
            @click.native="toTarget(card.name)"
          >
            <el-card shadow="hover" class="grid-content">
              <i :class="card.icon" :style="{ color: card.color }" />
              <p>{{ card.label }}</p>
            </el-card>
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="12" class="bottom">
        <el-row :gutter="32">
          <el-col :xs="36" :sm="36" :lg="20">
            <div class="chart-player">
              <todo-list />
            </div>
          </el-col>
        </el-row>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import TodoList from './component/todoList'
import { mapGetters } from 'vuex'
export default {
  name: 'Dashboard',
  components: {
    TodoList // TodoList
  },
  data() {
    return {
      toolCards: [
        {
          label: '用户管理',
          icon: 'el-icon el-icon-monitor',
          name: 'user',
          color: '#ff9c6e'
        },
        {
          label: '角色管理',
          icon: 'el-icon el-icon-setting',
          name: 'authority',
          color: '#69c0ff'
        },
        {
          label: '菜单管理',
          icon: 'el-icon el-icon-menu',
          name: 'menu',
          color: '#b37feb'
        },
        {
          label: '操作历史',
          icon: 'el-icon el-icon-time',
          name: 'operation',
          color: '#ffd666'
        },
        {
          label: '系统配置',
          icon: 'el-icon el-icon-s-operation',
          name: 'system',
          color: '#ff85c0'
        },
        {
          label: '服务器状态',
          icon: 'el-icon el-icon-cloudy',
          name: 'state',
          color: '#5cdbd3'
        }
      ]
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name })
    }
  }
}
</script>

<style lang="scss" scoped>
.big {
  margin: 100px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;
  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;
    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }
  .mid {
    width: 100%;
    height: 380px;
    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }
    .shadow {
      // width: 50%;
      .el-col {
        margin-bottom: 20px;
        .grid-content {
          padding: 36px 0;
        }
      }
    }
    .bottom {
      // width: 50%;
      height: 300px;
      // margin: 20px 0;
      .el-row {
        margin-right: 4px !important;
      }
      .chart-player {
        margin-left: 60px;
        width: 100%;
        height: 100%;
        padding: 10px;
        background-color: #fff;
      }
    }
}
</style>
