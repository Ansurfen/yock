from setuptools import setup, find_packages

setup(
    name='yock',
    version='0.0.18',
    description="Yock is a solution of cross platform to compose distributed build stream.",
    author="ansurfen",
    author_email="axf593161@gmail.com",
    url="https://github.com/Ansurfen/yock",
    packages=['yock'],
    package_data={
        'yock': ['*.*', '**/*.*']
    },
    keywords="package-manager, distributed-systems, build-tool, lua",
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'Topic :: Software Development :: Build Tools',
        'License :: OSI Approved :: MIT License'
    ]
)
