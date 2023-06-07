import { Balance, Button, Input } from '@massalabs/react-ui-kit';
import { FiArrowUpRight, FiPlus } from 'react-icons/fi';
import Intl from '../../../i18n/i18n';
import Modal from './Advanced';
import ContactList from './ContactList';

export function SendForm({ ...props }) {
  const {
    amount,
    account,
    formattedBalance,
    recipient,
    error,
    fees,
    modal,
    setModal,
    modalAccounts,
    setModalAccounts,
    handleModalAccounts,
    handleFees,
    setRecipient,
    handleSubmit,
    handleChange,
    SendPercentage,
  } = props;

  const modalArgsAdvanced = {
    fees,
    modal,
    setModal,
    handleFees,
  };
  const modalArgsAccounts = {
    modalAccounts,
    setModalAccounts,
    handleModalAccounts,
    setRecipient,
    account,
  };

  return (
    <form onSubmit={handleSubmit}>
      {/* Balance Section */}
      <div>
        <p> {Intl.t('sendcoins.account-balance')}</p>
        <Balance customClass="pl-0 bg-transparent" amount={formattedBalance} />
      </div>
      <div className="flex flex-row justify-between w-full pb-3.5 ">
        <p> {Intl.t('sendcoins.send-action')} </p>
        <p>
          {Intl.t('sendcoins.available-balance')} <u>{formattedBalance}</u>
        </p>
      </div>
      <div className="pb-3.5">
        <Input
          placeholder={'Amount to send'}
          value={amount}
          name="amount"
          onChange={(e) => handleChange(e)}
          error={error?.amount}
        />
      </div>
      <div className="flex flex-row-reverse">
        <ul className="flex flex-row">
          <SendPercentage percentage={25} />
          <SendPercentage percentage={50} />
          <SendPercentage percentage={75} />
          <SendPercentage percentage={100} />
        </ul>
      </div>
      <p className="pb-3.5">{Intl.t('sendcoins.recipient')}</p>
      <div className="pb-3.5">
        <Input
          placeholder={'Recipient'}
          value={recipient}
          onChange={(e) => setRecipient(e.target.value)}
          name="recipient"
          error={error?.recipient}
        />
      </div>
      <div className="flex flex-row-reverse pb-3.5">
        <p
          className="hover:cursor-pointer"
          onClick={() => setModalAccounts(!modalAccounts)}
        >
          <u>{Intl.t('sendcoins.transfer-between-acc')}</u>
        </p>
      </div>
      {/* Button Section */}
      <div className="flex flex-col w-full gap-3.5">
        <Button
          onClick={() => setModal(!modal)}
          variant={'secondary'}
          posIcon={<FiPlus />}
        >
          {Intl.t('sendcoins.advanced')}
        </Button>

        <div>
          <Button type="submit" posIcon={<FiArrowUpRight />}>
            {Intl.t('sendcoins.send')}
          </Button>
        </div>
      </div>
      <div>
        {modal ? (
          <Modal {...modalArgsAdvanced} />
        ) : modalAccounts ? (
          <ContactList {...modalArgsAccounts} />
        ) : null}
      </div>
    </form>
  );
}