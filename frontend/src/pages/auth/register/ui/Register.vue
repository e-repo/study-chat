<template>

  <v-container class="h-100">
    <v-row class="h-100 align-center">
      <v-col class="d-flex justify-center">

        <v-card
            min-width="340"
            width="400"
            elevation="4"
        >
          <v-card-item>
            <v-card-title>
              <h4 class="text-center">Регистрация</h4>
            </v-card-title>
            <v-alert
                v-if="createUserForm.serverError"
                text="Не корректные логин или пароль, проверьте правильность ввода учетных данных!"
                type="error"
                variant="outlined"
            ></v-alert>
            <v-alert
                v-if="createUserForm.isCreated"
                text="Вы успешно зарегистрированны! Перейдите на форму входа для заполнения учетных данных, или дождитесь автоматического перехода через 5сек."
                type="error"
                variant="outlined"
            ></v-alert>
          </v-card-item>

          <v-divider class="mx-4 mb-2"></v-divider>

          <v-card-text>

            <v-form
                v-model="createUserForm.isValid"
                @submit.prevent="onSubmit"
            >
              <v-text-field
                  v-model="createUserForm.firstName"
                  :rules="[firstNameRules.required, firstNameRules.counter]"
                  label="Имя пользователя"
                  variant="underlined"
                  counter
                  @click:append-inner="isPassShow = !isPassShow"
              ></v-text-field>

              <v-text-field
                  v-model="createUserForm.email"
                  :rules="[emailRules.required, emailRules.email]"
                  type="email"
                  label="Email"
                  variant="underlined"
              ></v-text-field>

              <v-text-field
                  v-model="createUserForm.password"
                  :rules="[passRules.required, passRules.counter]"
                  :type="isPassShow ? 'password' : 'text'"
                  label="Пароль"
                  variant="underlined"
                  :append-inner-icon="isPassShow ? 'mdi-eye-off' : 'mdi-eye'"
                  counter
                  @click:append-inner="isPassShow = !isPassShow"
              ></v-text-field>

              <div
                  class="d-flex flex-wrap justify-end mt-4"
              >
                <v-btn
                    :disabled="!createUserForm.isValid"
                    :loading="createUserForm.loading"
                    class="w-100"
                    type="submit"
                    color="success"
                >
                  Отправить
                </v-btn>
              </div>

            </v-form>

          </v-card-text>

          <v-divider class="mx-4"></v-divider>

          <v-card-actions
              class="ma-2 flex-column align-end"
          >
            <p class="mt-2">
              <router-link to="/login" class="text-info">Вход</router-link>
            </p>
          </v-card-actions>
        </v-card>

      </v-col>
    </v-row>
  </v-container>

</template>

<script setup lang="ts">
import { FormHelper } from '@/shared/lib';
import { reactive, ref, watch } from 'vue';
import {CreateUser, useUserModel} from '@/entities/user';
import { AxiosError } from 'axios';
import { useRouter } from 'vue-router';

const userModel = useUserModel();
const router = useRouter();

const isPassShow = ref<boolean>(true);

interface CreateUserForm {
  serverError: boolean;
  isValid: boolean;
  firstName: string | null
  email: string | null;
  password: string | null;
  loading: boolean;
  isCreated: boolean;
}

const createUserForm = reactive<CreateUserForm>({
  serverError: false,
  isValid: false,
  firstName: null,
  email: null,
  password: null,
  loading: false,
  isCreated: false
})

const emailRules = {
  required: FormHelper.requiredRule,
  email: (value: string): FormHelper.RuleType => FormHelper.emailPattern.test(value) || 'Некорректный \'Email\'',
};

const passRules = {
  required: FormHelper.requiredRule,
  counter: (value: string): FormHelper.RuleType => {
    if (value.length < 4) {
      return 'Длинна пароля не менее 4-ти символов';
    }

    return value.length <= 20 || 'Максимальное число символов 20';
  }
};

const firstNameRules = {
  required: FormHelper.requiredRule,
  counter: (value: string): FormHelper.RuleType => {
    if (value.length < 3) {
      return 'Длинна имени не менее 3-eх символов';
    }

    return value.length <= 50 || 'Максимальное число символов 50';
  }
};

const onSubmit = async (): Promise<void> => {
  if (! createUserForm.isValid) {
    return;
  }

  createUserForm.loading = true;

  try {
    await userModel.signUp(<CreateUser>{
      firstName: createUserForm.firstName,
      email: createUserForm.email,
      password: createUserForm.password,
    });

    createUserForm.email = null;
    createUserForm.password = null;
    createUserForm.firstName = null;

    createUserForm.loading = false;
    createUserForm.isCreated = true;

    setTimeout(async () => {
      await router.push({ name: 'Login' });
    }, 5000);
  } catch (error: any) {
    createUserForm.loading = false;
    createUserForm.isValid = false;

    if (error instanceof AxiosError) {
      createUserForm.serverError = true;

      return;
    }
  }
};

watch(
    [() => createUserForm.firstName, () => createUserForm.email, () => createUserForm.password],
    (): void => { createUserForm.serverError = false}
)
</script>

<style scoped></style>
