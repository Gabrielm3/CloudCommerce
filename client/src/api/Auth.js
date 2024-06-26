import { Auth } from "@aws-amplify/auth";

async function register(email, password) {
  try {
    const response = await Auth.signUp({
      username: email,
      password,
    });
    return response;
  } catch (error) {
    throw error;
  }
}

async function resendCode(email) {
  try {
    await Auth.resendSignUp(email);
  } catch (error) {
    throw error;
  }
}

async function confirmation(email, code) {
  try {
    await Auth.confirmSignUp(email, code);
    return true;
  } catch (error) {
    throw error;
  }
}

async function login(email, password) {
  try {
    await Auth.signIn({
      username: email,
      password,
    });

    const session = await Auth.currentAuthenticatedUser({
      bypassCache: false,
    });
    return session;
  } catch (error) {
    throw error;
  }
}

export const authApi = {
  register,
  resendCode,
  confirmation,
  login,
};
