import React, { useEffect, useState } from "react";
import ReactDOM from "react-dom";
import { Select } from "antd";
import * as ls from "local-storage";
import { searchTopicRequest as searchRequest } from "../service";
const { Option, OptGroup } = Select;
let timeout;
let currentValue;
const MAX_LIMIT = 3;
function searchTopic(value, callback) {
  if (timeout) {
    console.log("clearTimeout");
    clearTimeout(timeout);
    timeout = null;
  }
  currentValue = value;

  function search() {
    searchRequest(value).then((d) => {
      if (currentValue === value) {
        const { data } = d;
        const result = [];
        data.forEach((r) => {
          result.push({
            value: r.id,
            text: r.name,
          });
        });
        callback(result);
      }
    });
  }

  timeout = setTimeout(search, 500);
}

function SearchInput(props) {
  const { onChange } = props;
  const [data, setData] = useState([]);
  const [value, setValue] = useState();
  const [used, setUsed] = useState([]);
  const handleSearch = (value) => {
    if (value) {
      searchTopic(value, (data) => setData(data));
    } else {
      setData([]);
    }
  };

  const handleChange = (val, option) => {
    if (value == val) {
      return;
    }
    setValue(val);
    onChange(val);
    let filted = used.filter((p) => p.value != val);
    let newUsed = [{ value: val, text: option.children }, ...filted];
    newUsed = newUsed.slice(0, MAX_LIMIT);
    setUsed(newUsed);
    ls.set("TOPIC_USED", newUsed);
  };

  useEffect(() => {
    let lsdata = ls.get("TOPIC_USED");
    console.log("init used");
    if (Array.isArray(lsdata)) {
      if (lsdata.length > 0) {
        setUsed(lsdata);
        setValue(lsdata[0].value);
        onChange(lsdata[0].value);
      }
      return;
    }

    //保存
    return () => {
      console.log("save used");
    };
  }, []);

  const options = data.map((d) => (
    <Option key={"u_" + d.value} value={d.value}>
      {d.text}
    </Option>
  ));
  const usedOptions = used.map((d) => (
    <Option key={"s_" + d.value} value={d.value}>
      {d.text}
    </Option>
  ));
  return (
    <Select
      showSearch
      value={value}
      placeholder={props.placeholder}
      style={props.style}
      defaultActiveFirstOption={false}
      showArrow={false}
      filterOption={false}
      onSearch={handleSearch}
      onChange={handleChange}
      notFoundContent={null}
    >
      <OptGroup label="最近使用" key="used">
        {usedOptions}
      </OptGroup>
      <OptGroup label="输入搜索" key="search">
        {options}
      </OptGroup>
    </Select>
  );
}

export default SearchInput;
