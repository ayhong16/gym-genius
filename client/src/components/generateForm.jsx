import TimePicker from "./timePicker";


export default function GenerateForm() {
  return (
    <div className="flex flex-col">
      <h1 className="text-left text-2xl mb-4">Generate Workout</h1>
      <h3 className="text-left">Name:</h3>
      <input type="text" className="border-2 border-black rounded-md p-1 mb-4" />
      <div className="flex flex-row">
        <h3 className="text-left">Length:</h3>
        <TimePicker />
      </div>
    </div>
  )
}