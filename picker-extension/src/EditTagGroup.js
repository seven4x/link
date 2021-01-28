import React, { useState, useRef,useEffect } from "react";
import ReactDOM from "react-dom";
import { Tag, Input, Tooltip } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import "./EditTagGroup.css";

function EditableTagGroup(props) {
  const {onChange} = props
  const [state, setState] = useState({
    tags: [],
    inputVisible: false,
    inputValue: "",
    editInputIndex: -1,
    editInputValue: "",
  });

  const inputRef = useRef(null);
  const editInputRef = useRef(null);
  useEffect(()=>{
    if (inputRef.current != null ){
      inputRef.current.focus();
    }
  },[state.inputVisible])

  useEffect(()=>{
    if (editInputRef.current != null ){
      editInputRef.current.focus();
    }
  },[state.editInputIndex])

  useEffect(()=>{
    onChange(state.tags)
  },[state.tags])

  const handleClose = (removedTag) => {
    const tags = state.tags.filter((tag) => tag !== removedTag);
     
    setState({ ...state, tags });
  };

  const showInput = () => {
    setState({ ...state, inputVisible: true });
   
   
  };

  const handleInputChange = (e) => {
    setState({ ...state, inputValue: e.target.value });
  };

  const handleInputConfirm = () => {
    const { inputValue } = state;
    let { tags } = state;
    if (inputValue && tags.indexOf(inputValue) === -1) {
      tags = [...tags, inputValue];
    }
   
    setState({
      ...state,
      tags,
      inputVisible: false,
      inputValue: "",
    });
  };

  const handleEditInputChange = (e) => {
    setState({ ...state, editInputValue: e.target.value });
  };

  const handleEditInputConfirm = () => {
    let { tags, editInputIndex, editInputValue } = state;
    const newTags = [...tags];
    newTags[editInputIndex] = editInputValue;
    setState({
      ...state,
      tags: newTags,
      editInputIndex: -1,
      editInputValue: "",
    });
  };

 
  const {
    tags,
    inputVisible,
    inputValue,
    editInputIndex,
    editInputValue,
  } = state;

  return (
    <>
      {tags.map((tag, index) => {
        if (editInputIndex === index) {
          return (
            <Input
              ref={editInputRef}
              key={tag}
              size="small"
              className="tag-input"
              value={editInputValue}
              onChange={handleEditInputChange}
              onBlur={handleEditInputConfirm}
              onPressEnter={handleEditInputConfirm}
            />
          );
        }

        const isLongTag = tag.length > 20;

        const tagElem = (
          <Tag
            className="edit-tag"
            key={tag}
            closable={true}
            onClose={() => handleClose(tag)}
          >
            <span
              onDoubleClick={(e) => {
                setState({
                  ...state,
                  editInputIndex: index,
                  editInputValue: tag,
                });
                e.preventDefault();
              }}
            >
              {isLongTag ? `${tag.slice(0, 20)}...` : tag}
            </span>
          </Tag>
        );
        return isLongTag ? (
          <Tooltip title={tag} key={tag}>
            {tagElem}
          </Tooltip>
        ) : (
          tagElem
        );
      })}
      {inputVisible && (
        <Input
          ref={inputRef}
          type="text"
          size="small"
          className="tag-input"
          value={inputValue}
          onChange={handleInputChange}
          onBlur={handleInputConfirm}
          onPressEnter={handleInputConfirm}
        />
      )}
      {!inputVisible && (
        <Tag className="site-tag-plus" onClick={showInput}>
          <PlusOutlined /> 添加标签
        </Tag>
      )}
    </>
  );
}

export default EditableTagGroup;
