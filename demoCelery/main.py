from celery import Celery

app = Celery('tasks',
    broker='redis://localhost:6379/1',
    backend='redis://localhost:6379/2'
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
    每5s调度一次minus
    """
    print('-> sender:{}'.format(sender))
    sender.add_periodic_task(5.0, minus.s())

@app.task
def minus():
    """
    1亿减到1.
    cpu计算型任务.
    """
    x = 10000000
    # while x > 1:
    #     x -= 1
    print("x ==> done")

## celery -A main beat
