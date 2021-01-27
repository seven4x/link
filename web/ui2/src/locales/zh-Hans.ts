import messages from './zh-Hans.messages.json';


export default {
  // 合并所有 messages, 加入组件的 messages
  messages: messages,

  // locale
  code: 'zh-CN',


  // 自定义 formates
  formats: {
    // 日期、时间
    date: {
      normal: {
        hour12: false,
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
      },
    },
    // 货币
    money: {
      currency: 'CNY',
    },
  },
};
