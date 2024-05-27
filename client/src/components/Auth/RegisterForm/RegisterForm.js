import { Button, Form } from "semantic-ui-react";
import Link from "next/link";
import { useFormik } from "formik";
import { initialValues, validationSchema } from "./RegisterForm.form";
import styles from "./RegisterForm.module.scss";

export function RegisterForm() {
  const formik = useFormik({
    initialValues: initialValues(),
    validationSchema: validationSchema(),
    validationOnChange: false,
    onSubmit: (values) => {
      console.log(values);
    },
  });

  return (
    <>
      <Form onSubmit={formik.handleSubmit}>
        <Form.Input
          name="email"
          placeholder="email"
          value={formik.values.email}
          onChange={formik.handleChange}
          error={formik.errors.email}
        />
        <Form.Input
          type="password"
          name="password"
          placeholder="password"
          value={formik.values.password}
        />
        <Form.Input
          type="password"
          name="repeatpassword"
          placeholder="repeat password"
        />
        <Form.Button type="submit" fluid loading={formik.isSubmitting}>
          Register
        </Form.Button>
      </Form>

      <p className={styles.login}>Already have an account</p>
      <Button as={Link} href="/join/login" fluid basic>
        Login
      </Button>
    </>
  );
}
