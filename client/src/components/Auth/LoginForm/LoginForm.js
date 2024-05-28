import { configurationValues } from "../ConfirmationForm/ConfirmationForm.form";
import { Form, Button } from "semantic-ui-react";
import { useFormik } from "formik";
import { useAuth } from "@/hooks";
import { configurationValues, validationSchema } from "./LoginForm.form";
import styles from "./LoginForm.module.css";
import { authApi } from "@/api";

export function LoginForm() {
  const { login } = useAuth();

  const formik = useFormik({
    initialValues: configurationValues(),
    validationSchema: validationSchema(),
    validateOnChange: false,
    onSubmit: async (formValue) => {
      try {
        await authApi.login(formValue.email, formValue.password);
        await login();
      } catch (error) {
        console.error(error);
      }
    },
  });

  return (
    <Form onSubmit={formik.handleSubmit}>
      <Form.Input
        name="email"
        placeholder="Email"
        value={formik.values.email}
        onChange={formik.handleChange}
        error={formik.errors.email}
      />
      <Form.Input
        type="password"
        name="password"
        placeholder="Password"
        value={formik.values.password}
        onChange={formik.handleChange}
        error={formik.errors.password}
      />
      <Form.Button type="submit" fluid loading={formik.isSubmitting}>
        Login
      </Form.Button>

      <p className={styles.register}>New client</p>
      <Button as={Link} href="/join/register" fluid basic>
        Register
      </Button>
    </Form>
  );
}
