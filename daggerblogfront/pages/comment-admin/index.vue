<template>
  <div class="comment-admin-page">
      <div style="width:100%;height:100%;">
          <app-header></app-header>
          <div class="table-container">
              <div v-if="totalCount>0" class="pagi-container">
               <el-pagination style="margin: 0 auto" class="pagination" layout="prev, pager, next" :total="totalCount" :page-size="pageSize"
               v-on:current-change="changePage">
                </el-pagination>
        </div>
              <el-table
          :border=true
          :data="comment_list"
          style="width: 100%">
          <el-table-column
            label="评论时间"
            prop="comment_time"
            width="120">
          </el-table-column>
          <el-table-column
            label="评论者"
            prop="commenter"
            width="120">
          </el-table-column>
          <el-table-column
            label="文章标题"
            prop="post_title"
            width="170">
          </el-table-column>
          <el-table-column
            label="评论内容"
            prop="comment_content"
            width="450">
          </el-table-column>
     
          <el-table-column
            label="操作"
            width="140">
            <template slot-scope="scope">
              <div style="display:flex;">
                <el-button size="small" type="success" @click="goPage('/post/'+scope.row.post_id)">前往</el-button>
                <el-button size="small" type="danger" @click="deleteComment(scope.row.id)">删除</el-button>
              </div>
            </template>
          </el-table-column>

        </el-table>
        <div v-if="totalCount>0" class="pagi-container">
               <el-pagination style="margin: 0 auto" class="pagination" layout="prev, pager, next" :total="totalCount" :page-size="pageSize"
               v-on:current-change="changePage">
                </el-pagination>
        </div>
          </div>
      </div>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
import base from '~/mixins/base'
import {AXIOS} from '~/common/http-commons'
export default {
  components: {
      'app-header': Header
  },
  mixins:[base],
  data(){
    return{
        totalCount: 0,
        currentPage: 1,
        pageSize: 0,
        comment_list:[]
    }
  },
  methods:{
      deleteComment(id){
          this.$confirm('确认删除该评论？')
      .then(_ => {
          AXIOS.delete('/api/v1/comment/'+id).then(response=>{
              if(response.data.ok==true){
                  this.$notify({
                          title: '成功',
                          type: 'success',
                          message: response.data.message
                        });
                  this.loadMyComments();
              }else{
                  this.$notify.error({
                          title: '错误',
                          message: response.data.message
                        });
              }
          }).catch(e=>{
               this.$notify.error({
                          title: '错误',
                          message: "未知错误"
                        });
          })
          done();
      }).catch(_ => {});

          
      },
      changePage: function (currentPage) {
       this.currentPage = currentPage;
       this.loadMyComments();
      },
      loadMyComments(){
          var params = {};
          params.page = this.currentPage;
          AXIOS.get('/api/v1/comment/list',{
            params: params
          }).then(response=>{
              if(response.data.ok==true){
                  this.comment_list = response.data.data.data;
                  this.totalCount = Number(response.data.data.totalCount);
                  this.currentPage = Number(response.data.data.currentPage);
                  this.pageSize = Number(response.data.data.pageSize);

              }else{
                  this.$notify.error({
                          title: '错误',
                          message: response.data.message
                        });
              }
          }).catch(e=>{
                 this.$notify.error({
                          title: '错误',
                          message: '载入未读评论未知错误'
                        });
          })
      }
  },
  mounted(){
      this.loadMyComments();
  }
}
</script>

<style scoped>
.comment-admin-page {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
.table-container{
    width: 1000px;
    min-height:200px;
    margin: 0 auto;
    margin-top:30px;
}

</style>

