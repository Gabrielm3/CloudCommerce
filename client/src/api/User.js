import { ENV } from "@/utils";
import { Auth } from "@aws-amplify/auth";

async function me() {
  try {
    const session = await Auth.currentSession();
    const token = session.acessToken.jwtToken();
    console.log(session);

    const url = `${ENV.API_URL}/${ENV.ENDPOINTS.USER_ME}`;
    const params = {
      headers: {
        Authorization: token,
      },
    };

    const response = await fetch(url, params);
    const result = await response.json();

    if (response.status !== 200) throw result;

    return result;
  } catch (error) {
    throw error;
  }
}

export const userApi = {
  me,
};
