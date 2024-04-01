import asyncio
import logging
import os

from aiogram import Bot, Dispatcher, types
from aiogram.filters.command import Command
from aiogram.utils.keyboard import ReplyKeyboardBuilder

# Включаем логирование, чтобы не пропустить важные сообщения
logging.basicConfig(level=logging.INFO)
# Объект бота
bot = Bot(token=os.getenv('WORKER_BOT_TOKEN'), parse_mode="HTML")
# Диспетчер
dp = Dispatcher()


@dp.message(Command("special_buttons"))
async def cmd_special_buttons(message: types.Message):
    builder = ReplyKeyboardBuilder()
    # метод row позволяет явным образом сформировать ряд
    # из одной или нескольких кнопок. Например, первый ряд
    # будет состоять из двух кнопок...
    builder.row(
        types.KeyboardButton(text="Запросить геолокацию", request_location=True),
        types.KeyboardButton(text="Запросить контакт", request_contact=True))
    await message.answer("Выберите действие:", reply_markup=builder.as_markup(resize_keyboard=True))


# Хэндлер на команду /start
@dp.message(Command("start"))
async def cmd_start(message: types.Message):
    await message.answer(f'Здравствуйте! Вы зашли в ')


@dp.message()
async def echo(msg: types.Message):
    print(msg.from_user.full_name)
    print(msg.from_user.id)
    print(msg.text)
    await bot.send_message(text=msg.text, chat_id=msg.from_user.id, parse_mode="Markdown")


# Запуск процесса поллинга новых апдейтов
async def main():
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())