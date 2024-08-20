export default function TimePicker() {
  return (
    <div className="flex flex-row items-center mx-1">
      <select id={"hourDropdown"} className="preline-dropdow" value={0}>
        {[...Array(24).keys()].map((hour) => (
          <option key={hour} value={hour.toString()}>
            {hour}
          </option>
        ))}
      </select>
      <label htmlFor={"minuteDropdown"} className="mr-1">hours</label>
      <select id={"minuteDropdown"} className="preline-dropdow" value={0}>
        {[...Array(60).keys()].map((hour) => (
          <option key={hour} value={hour.toString()}>
            {hour}
          </option>
        ))}
      </select>
      <label htmlFor={"minuteDropdown"}>mins</label>
    </div>
  );
}