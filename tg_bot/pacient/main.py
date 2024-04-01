import asyncio
import logging
import os

import dotenv
from aiogram import Bot, Dispatcher, types, F
from aiogram.filters.command import Command
from aiogram.utils import keyboard

from pacient.static import patient_messages

dotenv.load_dotenv()

logging.basicConfig(level=logging.INFO)
bot = Bot(token=os.getenv('PATIENT_BOT_TOKEN'), parse_mode="HTML")
dp = Dispatcher()


@dp.message(Command("start"))
async def cmd_start(patient_start_message: types.Message):
    patient_auth = False
    builder = keyboard.ReplyKeyboardBuilder()
    if not patient_auth:
        builder.row(types.KeyboardButton(text="Войти", request_contact=True))
        await patient_start_message.answer(patient_messages.initial_message(),
                                           reply_markup=builder.as_markup(resize_keyboard=True))


@dp.message(F.content_type == types.ContentType.TEXT)
async def message(patient_unknown_message: types.Message):
    print('patient_unknown_message')
    print(patient_unknown_message)


@dp.message(F.content_type == types.ContentType.CONTACT)
async def contact(patient_contact_message: types.Contact):
    print('patient_contact_message')
    print(patient_contact_message.contact.phone_number)
    patient_contact_is_exist = False
    builder = keyboard.KeyboardBuilder()
    if not patient_contact_is_exist:
        builder.row(types.KeyboardButton(text="Изменить номер здесь"), types.KeyboardButton(text="Перейти на сайт"))
        builder.row(types.KeyboardButton(text="<b>Вызвать скорую помощь</b>"))
        await patient_contact_message.answer(patient_messages.unknown_phonenumber(),
                                             reply_markup=builder.as_markup(resize_keyboard=True))


async def main():
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())
