<template>
  <DashboardPage
    db-name="perfint"
    table="ideaSharedIndices"
    persistent-id="shared_indexes_dashboard"
    initial-machine="Linux EC2 C6i.8xlarge (32 vCPU Xeon, 64 GB)"
  >
    <section>
      <GroupProjectsChart
        label="Indexing (IntelliJ project)"
        measure="indexing"
        :projects="['ijx-intellij-speed/downloading-shared-indexes',
                    'ijx-intellij-speed/shared-indexes',
                    'ijx-intellij-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Scanning (IntelliJ project)"
        measure="scanning"
        :projects="['ijx-intellij-speed/downloading-shared-indexes',
                    'ijx-intellij-speed/shared-indexes',
                    'ijx-intellij-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing runs (IntelliJ project)"
        measure="numberOfIndexingRuns"
        :projects="['ijx-intellij-speed/downloading-shared-indexes',
                    'ijx-intellij-speed/shared-indexes',
                    'ijx-intellij-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by usual indexes (IntelliJ project)"
        measure="numberOfIndexedFiles"
        :projects="['ijx-intellij-speed/downloading-shared-indexes',
                    'ijx-intellij-speed/shared-indexes',
                    'ijx-intellij-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by shared indexes (IntelliJ project)"
        measure="numberOfFilesIndexedByExtensions"
        :projects="['ijx-intellij-speed/downloading-shared-indexes',
                    'ijx-intellij-speed/shared-indexes',
                    'ijx-intellij-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Updating time (IntelliJ project)"
        measure="updatingTime"
        :projects="['ijx-intellij-speed/downloading-shared-indexes',
                    'ijx-intellij-speed/shared-indexes',
                    'ijx-intellij-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Indexing (Community project)"
        measure="indexing"
        :projects="['ijx-intellij-community-speed/downloading-shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-community-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-community-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Scanning (Community project)"
        measure="scanning"
        :projects="['ijx-intellij-community-speed/downloading-shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-community-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-community-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing runs (Community project)"
        measure="numberOfIndexingRuns"
        :projects="['ijx-intellij-community-speed/downloading-shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-community-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-community-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by usual indexes (Community project)"
        measure="numberOfIndexedFiles"
        :projects="['ijx-intellij-community-speed/downloading-shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-community-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-community-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by shared indexes (Community project)"
        measure="numberOfFilesIndexedByExtensions"
        :projects="['ijx-intellij-community-speed/downloading-shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-community-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-community-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Updating time (Community project)"
        measure="updatingTime"
        :projects="['ijx-intellij-community-speed/downloading-shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes',
                    'ijx-intellij-community-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-intellij-community-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-intellij-community-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Indexing (JDK Maven project)"
        measure="indexing"
        :projects="['ijx-jdk-maven-speed/downloading-shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-maven-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-maven-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Scanning (JDK Maven project)"
        measure="scanning"
        :projects="['ijx-jdk-maven-speed/downloading-shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-maven-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-maven-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing runs (JDK Maven project)"
        measure="numberOfIndexingRuns"
        :projects="['ijx-jdk-maven-speed/downloading-shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-maven-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-maven-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by usual indexes (JDK Maven project)"
        measure="numberOfIndexedFiles"
        :projects="['ijx-jdk-maven-speed/downloading-shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-maven-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-maven-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by shared indexes (JDK Maven project)"
        measure="numberOfFilesIndexedByExtensions"
        :projects="['ijx-jdk-maven-speed/downloading-shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-maven-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-maven-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Updating time (JDK Maven project)"
        measure="updatingTime"
        :projects="['ijx-jdk-maven-speed/downloading-shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes',
                    'ijx-jdk-maven-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-maven-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-maven-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Indexing (JDK Gradle project)"
        measure="indexing"
        :projects="['ijx-jdk-gradle-speed/downloading-shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-gradle-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Scanning (JDK Gradle project)"
        measure="scanning"
        :projects="['ijx-jdk-gradle-speed/downloading-shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-gradle-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing runs (JDK Gradle project)"
        measure="numberOfIndexingRuns"
        :projects="['ijx-jdk-gradle-speed/downloading-shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-gradle-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by usual indexes (JDK Gradle project)"
        measure="numberOfIndexedFiles"
        :projects="['ijx-jdk-gradle-speed/downloading-shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-gradle-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Number of indexing files by shared indexes (JDK Gradle project)"
        measure="numberOfFilesIndexedByExtensions"
        :projects="['ijx-jdk-gradle-speed/downloading-shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-gradle-speed/usual-indexes']"
      />
    </section>
    <section>
      <GroupProjectsChart
        label="Updating time (JDK Gradle project)"
        measure="updatingTime"
        :projects="['ijx-jdk-gradle-speed/downloading-shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-archive-and-git-hashes',
                    'ijx-jdk-gradle-speed/shared-indexes-with-hashes-without-javaScript-plugin',
                    'ijx-jdk-gradle-speed/usual-indexes']"
      />
    </section>
  </DashboardPage>
</template>

<script setup lang="ts">
import GroupProjectsChart from "../charts/GroupProjectsChart.vue"
import DashboardPage from "../common/DashboardPage.vue"

</script>