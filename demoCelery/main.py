from celery import Celery

app = Celery('tasks',
    broker='redis://localhost:6379/1',
    backend='redis://localhost:6379/1'
)
app.conf.update(
    CELERY_TASK_SERIALIZER='json',
    CELERY_ACCEPT_CONTENT=['json'],
    CELERY_RESULT_SERIALIZER='json',
    CELERY_ENABLE_UTC=True,
    CELERY_TASK_PROTOCOL=1,
)

@app.on_after_configure.connect
def setup_periodic_tasks(sender, **kwargs):
    """
    需要启动周期任务.
    每5s调度一次minus
    """
    print('-> sender:{}'.format(sender))
    sender.add_periodic_task(5.0, minus.s())

@app.task(name='main.minus')
def minus(x):
    """
    1亿减到1.
    cpu计算型任务.
    """
    # x = 10000000
    # while x > 1:
    #     x -= 1
    print("x ==> done")
    return x


@app.task(name='main.add')
def add(x, y):
    return x + y
## celery -A main beat

if __name__ == '__main__':
    print("minus:{}".format(minus))
    # print("add:{}".format(add))
    # ar = add.apply_async((5456, 2879), serializer='json')
    # print("ar:{}".format(ar))
    # print(ar.get())
    task = minus.apply_async((100000000,), serializer='json',  expires=10)
    print(f"task:{task} status:{task.status}")
    print(task.get())