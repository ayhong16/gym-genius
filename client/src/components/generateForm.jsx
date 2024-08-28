import TimePicker from "./timePicker";
import { useFormik } from "formik";


export default function GenerateForm() {
  const formik = useFormik({
    initialValues: {
      name: "",
      selectedHours: "",
      selectedMinutes: "",
    }
  })
  return (
    <form>
      <h1 className="text-left text-2xl mb-4">Generate Workout</h1>
      <div className="flex flex-col input-container">
        <h3 className="text-left">Name</h3>
        <input
          id="name"
          name="name"
          type="text"
          placeholder="Name your workout"
          onChange={formik.handleChange}
          value={formik.values.name}
          className="border-2 border-black rounded-md p-1 mb-4" />
        <div className="flex flex-row">
          <h3 className="text-left">Length:</h3>
          <TimePicker />
        </div>
      </div>
    </form>
  )
}