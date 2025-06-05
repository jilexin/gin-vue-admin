<template>
  <div>
    <warning-bar
      title="需要提前配置email配置文件，为防止不必要的垃圾邮件，在线体验功能不开放此功能体验。"
    />
    <div class="gva-form-box">
      <el-form
        ref="emailForm"
        label-position="right"
        label-width="80px"
        :model="form"
      >
        <el-form-item label="草稿标题">
          <el-input 
            v-model="form.title" 
            placeholder="请输入草稿标题（可选，默认使用邮件主题）"
          />
        </el-form-item>
        <el-form-item label="目标邮箱">
          <el-input 
            v-model="form.to" 
            placeholder="请输入收件人邮箱"
          />
        </el-form-item>
        <el-form-item label="邮件主题">
          <el-input 
            v-model="form.subject" 
            placeholder="请输入邮件主题"
          />
        </el-form-item>
        <el-form-item label="邮件内容">
          <el-input 
            v-model="form.body" 
            type="textarea" 
            :rows="8"
            placeholder="请输入邮件内容"
          />
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTestEmail">发送测试邮件</el-button>
          <el-button @click="sendEmail" type="primary">发送邮件</el-button>
          <el-button @click="saveDraft" type="success">保存草稿</el-button>
          <el-button @click="viewDrafts" type="info">查看草稿</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 草稿列表弹窗 -->
    <el-dialog
      v-model="draftsDialogVisible"
      title="我的草稿"
      width="80%"
      :before-close="handleClose"
    >
      <el-table :data="drafts" style="width: 100%">
        <el-table-column prop="title" label="草稿标题" width="200" />
        <el-table-column prop="to" label="收件人" width="200" />
        <el-table-column prop="subject" label="邮件主题" width="200" />
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column prop="updated_at" label="更新时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.updated_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="loadDraft(scope.row)">加载</el-button>
            <el-button size="small" type="danger" @click="deleteDraft(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <!-- 分页 -->
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :small="false"
        :disabled="false"
        :background="true"
        layout="total, sizes, prev, pager, next, jumper"
        :total="pagination.total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        style="margin-top: 20px; justify-content: center"
      />
    </el-dialog>
  </div>
</template>

<script setup>
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { 
    emailTest, 
    sendEmail as sendEmailApi,
    saveEmailDraft,
    getEmailDrafts,
    deleteEmailDraft
  } from '@/plugin/email/api/email.js'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { reactive, ref } from 'vue'

  defineOptions({
    name: 'Email'
  })

  const emailForm = ref(null)
  const form = reactive({
    title: '',
    to: '',
    subject: '',
    body: ''
  })

  // 草稿相关状态
  const draftsDialogVisible = ref(false)
  const drafts = ref([])
  const pagination = reactive({
    page: 1,
    pageSize: 10,
    total: 0
  })

  // 发送测试邮件
  const sendTestEmail = async () => {
    try {
      const res = await emailTest()
      if (res.code === 0) {
        ElMessage.success('测试邮件发送成功')
      }
    } catch (error) {
      ElMessage.error('发送失败: ' + error.message)
    }
  }

  // 发送邮件
  const sendEmail = async () => {
    if (!form.to || !form.subject || !form.body) {
      ElMessage.warning('请填写完整的邮件信息')
      return
    }
    
    try {
      const res = await sendEmailApi({
        to: form.to,
        subject: form.subject,
        body: form.body
      })
      if (res.code === 0) {
        ElMessage.success('邮件发送成功，请查收')
        // 发送成功后清空表单
        clearForm()
      }
    } catch (error) {
      ElMessage.error('发送失败: ' + error.message)
    }
  }

  // 保存草稿
  const saveDraft = async () => {
    if (!form.to && !form.subject && !form.body) {
      ElMessage.warning('请至少填写一项内容')
      return
    }
    
    try {
      const res = await saveEmailDraft({
        title: form.title,
        to: form.to,
        subject: form.subject,
        body: form.body
      })
      if (res.code === 0) {
        ElMessage.success('草稿保存成功')
      }
    } catch (error) {
      ElMessage.error('保存失败: ' + error.message)
    }
  }

  // 查看草稿
  const viewDrafts = async () => {
    draftsDialogVisible.value = true
    await loadDraftList()
  }

  // 加载草稿列表
  const loadDraftList = async () => {
    try {
      const res = await getEmailDrafts({
        page: pagination.page,
        pageSize: pagination.pageSize
      })
      if (res.code === 0) {
        drafts.value = res.data.list || []
        pagination.total = res.data.total || 0
      }
    } catch (error) {
      ElMessage.error('获取草稿列表失败: ' + error.message)
    }
  }

  // 加载草稿到编辑器
  const loadDraft = (draft) => {
    form.title = draft.title || ''
    form.to = draft.to || ''
    form.subject = draft.subject || ''
    form.body = draft.body || ''
    draftsDialogVisible.value = false
    ElMessage.success('草稿已加载到编辑器')
  }

  // 删除草稿
  const deleteDraft = async (id) => {
    try {
      await ElMessageBox.confirm('确定要删除这个草稿吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      
      const res = await deleteEmailDraft(id)
      if (res.code === 0) {
        ElMessage.success('删除成功')
        await loadDraftList() // 重新加载列表
      }
    } catch (error) {
      if (error !== 'cancel') {
        ElMessage.error('删除失败: ' + error.message)
      }
    }
  }

  // 清空表单
  const clearForm = () => {
    form.title = ''
    form.to = ''
    form.subject = ''
    form.body = ''
  }

  // 格式化日期
  const formatDate = (dateString) => {
    if (!dateString) return ''
    const date = new Date(dateString)
    return date.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    })
  }

  // 分页事件处理
  const handleSizeChange = (val) => {
    pagination.pageSize = val
    pagination.page = 1
    loadDraftList()
  }

  const handleCurrentChange = (val) => {
    pagination.page = val
    loadDraftList()
  }

  // 关闭弹窗
  const handleClose = () => {
    draftsDialogVisible.value = false
  }
</script>

<style scoped>
.gva-form-box {
  padding: 20px;
}
</style>
