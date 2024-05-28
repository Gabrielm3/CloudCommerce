import { useRouter } from "next/router";
import { Form, Button } from "semantic-ui-react";
import { useState, useEffect } from "react";
import { useFormik } from "formik";
import { authApi } from "@/api";
import { Separator } from "@/components/Shared";
import { configurationValues, validationSchema } from "./ConfirmationForm.form";

export function ConfirmationForm() {
  const router = useRouter();
  const { query } = router;
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    formik.setFieldValue("email", query.email);
  }, [query]);

  const formik = useFormik({
    initialValues: configurationValues(),
    validationSchema: validationSchema(),
    validateOnChange: false,
    onSubmit: async (formValue) => {
      try {
        await authApi.confirmation(formValue.email, formValue.code);
        router.push("/join/login");
      } catch (error) {
        console.error(error);
      }
    },
  });

  const onResendCode = async () => {
    formik.setFieldError("email", false);

    if (!formik.values.email) {
      formik.setFieldError("email", true);
      return;
    }

    setLoading(true);
    // send email
    await authApi.resendCode(formik.values.email);
    setLoading(false);

    console.log("Resend code");
  };

  return (
    <>
      <Form onSubmit={formik.handleSubmit}>
        <Form.Input
          name="email"
          placeholder="Email"
          value={formik.values.email}
          onChange={formik.handleChange}
          error={formik.errors.email}
        />
        <Form.Input
          name="code"
          placeholder="Code"
          values={formik.values.code}
          onChange={formik.handleChange}
          error={formik.errors.code}
        />
        <Form.Button type="submit" fluid loading={formik.isSubmitting}>
          Confirm user
        </Form.Button>
      </Form>

      <Separator height={50} />

      <Button fluid basic onClick={onResendCode} loading={loading}>
        Resend code
      </Button>
    </>
  );
}
